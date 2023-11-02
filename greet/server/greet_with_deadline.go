package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with: %v", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled, "The client canceled")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{Result: "Hellow " + in.FirstName}, nil
}
