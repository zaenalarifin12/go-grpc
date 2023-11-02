package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("doGreetWithDeadline was invoked\n")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Zainal"}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!!")
				return
			} else {
				log.Fatalf("Unexpected GRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non grpc error: %v\n", err)
		}
	}

	log.Printf("greet with deadline: %s\n", res.Result)
}
