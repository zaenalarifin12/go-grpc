package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{Number: 12390392840}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling names %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v\n", err)
		}

		log.Printf("Primes: %v", res.Result)
	}
}
