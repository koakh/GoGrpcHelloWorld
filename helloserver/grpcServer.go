package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"koakh.com/helloGrpc/hellopb"
	"google.golang.org/grpc"
)

type server struct{}

// receiver ais a pointer of *server
func (*server) Hello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := request.Name
	// get value from pointer &hellopb
	response := &hellopb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}

func main() {
	// register a server and let it listen on address 0.0.0.0:50051 , which is localhost and port 50051 is the port for gRPC connection
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()

	// The function RegisterHelloServiceServer is auto-generated from the protobuf file we just wrote.
	// If you look into the function, you would find that the second argument requires a struct type to implement the service interface
	hellopb.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}
