package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/calculator/proto"
	"log"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream %v\n", err)

	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending Number: %d\n", number)
		stream.Send(&pb.AvgRequest{Number: number})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
