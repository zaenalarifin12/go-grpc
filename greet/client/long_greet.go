package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	// List of unique names
	names := []string{"Zainal", "John", "Emily", "Michael", "Sara", "David", "Sophia", "James", "Olivia", "Daniel"}

	reqs := make([]*pb.GreetRequest, 0)

	// Create 100 instances of GreetRequest with unique names
	for i := 0; i < 100; i++ {
		nameIndex := i % len(names)
		req := &pb.GreetRequest{
			FirstName: names[nameIndex],
		}
		reqs = append(reqs, req)
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling longGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending Req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from longGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
