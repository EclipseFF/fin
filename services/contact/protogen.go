//go:build proto
// +build proto

package contact

//go:generate protoc --proto_path=. --proto_path=proto	--go_opt=paths=source_relative	--go_out=internal/delivery/grpc/interface contact.proto
//go:generate protoc --proto_path=. --proto_path=proto	--go-grpc_opt=paths=source_relative	--go-grpc_out=internal/delivery/grpc/interface contact.proto
