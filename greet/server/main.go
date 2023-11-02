package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
)

var addr string = "0.0.0.0:5000"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed To Listen: %v", err)
	}

	log.Printf("listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
