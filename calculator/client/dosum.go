package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		log.Fatalf("------ failed: %v", err)
	}
	log.Printf("Greeting: %s\n", res.Result)

}
