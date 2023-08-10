// Package pam implements the pam grpc service protocol to the daemon.
package pam

import (
	"context"
	"errors"
	"fmt"

	"github.com/ubuntu/authd"
	"github.com/ubuntu/authd/internal/brokers"
	"github.com/ubuntu/authd/internal/log"
	"github.com/ubuntu/decorate"
)

var _ authd.PAMServer = Service{}

// Service is the implementation of the PAM module service.
type Service struct {
	brokerManager *brokers.Manager

	authd.UnimplementedPAMServer
}

// NewService returns a new PAM GRPC service.
func NewService(ctx context.Context, brokerManager *brokers.Manager) Service {
	log.Debug(ctx, "Building new GRPC PAM service")

	return Service{
		brokerManager: brokerManager,
	}
}

// AvailableBrokers returns the list of all brokers with their details.
func (s Service) AvailableBrokers(ctx context.Context, _ *authd.Empty) (*authd.ABResponse, error) {
	var r authd.ABResponse

	for _, b := range s.brokerManager.AvailableBrokers() {
		r.BrokersInfos = append(r.BrokersInfos, &authd.ABResponse_BrokerInfo{
			Id:        b.ID,
			Name:      b.Name,
			BrandIcon: &b.BrandIconPath,
		})
	}

	return &r, nil
}

// GetPreviousBroker returns the previous broker set for a given user, if any.
func (s Service) GetPreviousBroker(ctx context.Context, req *authd.GPBRequest) (*authd.GPBResponse, error) {
	var r authd.GPBResponse

	b := s.brokerManager.BrokerForUser(req.GetUsername())
	if b != nil {
		r.PreviousBroker = &b.ID
	}

	return &r, nil
}

// SelectBroker starts a new session and selects the requested broker for the user.
func (s Service) SelectBroker(ctx context.Context, req *authd.SBRequest) (resp *authd.SBResponse, err error) {
	defer decorate.OnError(&err, "can't start authentication transaction")

	username := req.GetUsername()
	brokerID := req.GetBrokerId()
	lang := req.GetLang()

	if username == "" {
		return nil, errors.New("no user name provided")
	}
	if brokerID == "" {
		return nil, errors.New("no broker selected")
	}
	if lang == "" {
		lang = "C"
	}

	// Create a session and Memorize selected broker for it.
	sessionID, encryptionKey, err := s.brokerManager.NewSession(brokerID, username, lang)
	if err != nil {
		return nil, err
	}

	return &authd.SBResponse{
		SessionId:     sessionID,
		EncryptionKey: encryptionKey,
	}, err
}

// GetAuthenticationModes fetches a list of authentication modes supported by the broker depending on the session information.
func (s Service) GetAuthenticationModes(ctx context.Context, req *authd.GAMRequest) (resp *authd.GAMResponse, err error) {
	defer decorate.OnError(&err, "could not get authentication modes")

	sessionID := req.GetSessionId()
	if sessionID == "" {
		return nil, errors.New("no session ID provided")
	}

	broker, err := s.brokerManager.BrokerFromSessionID(sessionID)
	if err != nil {
		return nil, err
	}

	var layouts []map[string]string
	for _, l := range req.GetSupportedUiLayouts() {
		layout, err := uiLayoutToMap(l)
		if err != nil {
			return nil, err
		}
		layouts = append(layouts, layout)
	}

	authenticationModes, err := broker.GetAuthenticationModes(ctx, sessionID, layouts)
	if err != nil {
		return nil, err
	}

	var authModes []*authd.GAMResponse_AuthenticationMode
	for _, a := range authenticationModes {
		authModes = append(authModes, &authd.GAMResponse_AuthenticationMode{
			Id:    a["id"],
			Label: a["label"],
		})
	}

	return &authd.GAMResponse{
		AuthenticationModes: authModes,
	}, nil
}

// SelectAuthenticationMode set given authentication mode as selected for this sessionID to the broker.
func (s Service) SelectAuthenticationMode(ctx context.Context, req *authd.SAMRequest) (resp *authd.SAMResponse, err error) {
	defer decorate.OnError(&err, "can't select authentication mode")

	sessionID := req.GetSessionId()
	authenticationModeID := req.GetAuthenticationModeId()

	if sessionID == "" {
		return nil, errors.New("no session ID provided")
	}
	if authenticationModeID == "" {
		return nil, errors.New("no authentication mode provided")
	}

	broker, err := s.brokerManager.BrokerFromSessionID(sessionID)
	if err != nil {
		return nil, err
	}

	uiLayoutInfo, err := broker.SelectAuthenticationMode(ctx, sessionID, authenticationModeID)
	if err != nil {
		return nil, err
	}

	return &authd.SAMResponse{
		UiLayoutInfo: mapToUILayout(uiLayoutInfo),
	}, nil
}

// IsAuthorized returns broker answer to authorization request.
func (s Service) IsAuthorized(ctx context.Context, req *authd.IARequest) (resp *authd.IAResponse, err error) {
	defer decorate.OnError(&err, "can't check authorization")

	sessionID := req.GetSessionId()
	if sessionID == "" {
		return nil, errors.New("no session ID provided")
	}

	broker, err := s.brokerManager.BrokerFromSessionID(sessionID)
	if err != nil {
		return nil, err
	}

	access, data, err := broker.IsAuthorized(ctx, sessionID, req.GetAuthenticationData())
	if err != nil {
		return nil, err
	}

	return &authd.IAResponse{
		Access: access,
		Data:   data,
	}, nil
}

// SetDefaultBrokerForUser sets the default broker for the given user.
func (s Service) SetDefaultBrokerForUser(ctx context.Context, req *authd.SDBFURequest) (empty *authd.Empty, err error) {
	decorate.OnError(&err, "can't set default broker %q for user %q", req.GetBrokerId(), req.GetUsername())

	if req.GetUsername() == "" {
		return nil, errors.New("no user name given")
	}

	err = s.brokerManager.SetDefaultBrokerForUser(req.GetBrokerId(), req.GetUsername())

	return &authd.Empty{}, err
}

// EndSession asks the broker associated with the sessionID to end the session.
func (s Service) EndSession(ctx context.Context, req *authd.ESRequest) (empty *authd.Empty, err error) {
	decorate.OnError(&err, "could not abort session")

	sessionID := req.GetSessionId()
	if sessionID == "" {
		return nil, errors.New("no session id given")
	}

	return &authd.Empty{}, s.brokerManager.EndSession(sessionID)
}

func uiLayoutToMap(layout *authd.UILayout) (mapLayout map[string]string, err error) {
	if layout.GetType() == "" {
		return nil, fmt.Errorf("invalid layout option: type is required, got: %v", layout)
	}
	r := map[string]string{"type": layout.GetType()}
	if l := layout.GetLabel(); l != "" {
		r["label"] = l
	}
	if b := layout.GetButton(); b != "" {
		r["button"] = b
	}
	if w := layout.GetWait(); w != "" {
		r["wait"] = w
	}
	if e := layout.GetEntry(); e != "" {
		r["entry"] = e
	}
	if c := layout.GetContent(); c != "" {
		r["content"] = c
	}
	return r, nil
}

// mapToUILayout generates an UILayout from the input map.
func mapToUILayout(layout map[string]string) (r *authd.UILayout) {
	typ := layout["type"]
	label := layout["label"]
	entry := layout["entry"]
	button := layout["button"]
	wait := layout["wait"]
	content := layout["content"]

	return &authd.UILayout{
		Type:    typ,
		Label:   &label,
		Entry:   &entry,
		Button:  &button,
		Wait:    &wait,
		Content: &content,
	}
}
