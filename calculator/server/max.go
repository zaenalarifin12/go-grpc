package main

import (
	pb "github.com/zaenalarifin12/go-grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("max function was invoked")
	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{Result: maximum})

			if err != nil {
				log.Fatalf("Error while sending data %v", err)
			}
		}
	}
}
