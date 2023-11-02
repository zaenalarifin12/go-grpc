package main

import (
	"fmt"
	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreatManyTimes function was invoked with: %v\n ", in)

	for i := 0; i < 1000000000000; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}
