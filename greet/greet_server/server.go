package main

import (
	"fmt"
	"github.com/udodinho/grpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func main() {
	fmt.Println("Hello world!")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
