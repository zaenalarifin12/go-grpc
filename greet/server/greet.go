package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("greet function was invoked with %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.GetFirstName()}, nil
}
