// Package authd contains the autogenerated GRPC API between the modules and daemon.
package authd

//go:generate sh -c "PATH=\"$PATH:`go env GOPATH`/bin\" protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative authd.proto"
