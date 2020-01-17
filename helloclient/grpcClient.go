package main

import (
	"context"
	"fmt"
	"log"

	"koakh.com/helloGrpc/hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	// create a client hellopb.NewHelloServiceClient dialling to port 50051 and send the request through function client.Hello.
	// Note that all these functions are provided in the auto-generated file.
	client := hellopb.NewHelloServiceClient(cc)
	request := &hellopb.HelloRequest{Name: "Jeremy"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Receive response => [%v]\n", resp.Greeting)
}
