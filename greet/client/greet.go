package main

import (
	"context"
	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Zainal"})
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Printf("Greeting: %s\n", res.Result)

}
