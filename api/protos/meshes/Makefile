go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install storj.io/drpc/cmd/protoc-gen-go-drpc

protoc --go_out=. --go_opt=paths=source_relative --go-drpc_out=. --go-drpc_opt=paths=source_relative nexusops.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative nexusops.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto