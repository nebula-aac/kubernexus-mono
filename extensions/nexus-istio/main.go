package main

import grpcserver "github.com/nebula-aac/kubernexus-mono/pkg/grpc_server"

func main() {
	service := &grpcserver.Service{
		Port: "10000",
	}

	err := grpcserver.Start(service, nil)
	if err != nil {
		panic(err)
	}
}
