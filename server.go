package main

import (
	"context"
	"fmt"
	"github.com/rishabhgupta03/test/client/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	hello.UnimplementedHelloserviceServer
}

func (*server) Hello(ctx context.Context, request *hello.Hellorequest) (*hello.Helloresponse, error) {
	name := request.Name
	response := &hello.Helloresponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}
//func (*server) UnimplementedHelloserviceServer() {}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	hello.RegisterHelloserviceServer(s, &server{})

	s.Serve(lis)
}