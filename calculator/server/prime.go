package main

import (
	pb "github.com/zaenalarifin12/go-grpc/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("print function was invoked with %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{Result: divisor})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
