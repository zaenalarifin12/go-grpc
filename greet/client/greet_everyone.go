package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryOne(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{
			FirstName: "Zainal",
		},
		{
			FirstName: "Arifin",
		},
		{
			FirstName: "Arifin",
		},
		{
			FirstName: "Arifin",
		},
		{
			FirstName: "Arifin",
		}, {
			FirstName: "Arifin",
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send Request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
