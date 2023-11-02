package main

import (
	pb "github.com/zaenalarifin12/go-grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{Result: float64(sum) / float64(count)})
		}

		if err != nil {
			log.Fatalf("Error while reading stream:  %v\n", err)
		}

		log.Printf("receiving number: %v\n", req)
		sum += req.Number
		count++
	}
}
