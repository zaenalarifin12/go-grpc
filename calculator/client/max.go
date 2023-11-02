package main

import (
	"context"
	"fmt"
	pb "github.com/zaenalarifin12/go-grpc/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	fmt.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}

		for _, number := range numbers {
			log.Printf("Sending number: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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

				log.Fatalf("problem while reading server stream: %v", err)
				break
			}

			log.Printf("received a new maximum: %d\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
