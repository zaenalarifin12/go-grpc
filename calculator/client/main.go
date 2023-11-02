package main

import (
	pb "github.com/zaenalarifin12/go-grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:5000"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	if c == nil {
		log.Fatalf("Failed to create client")
	}
	//doAvg(c)
	doMax(c)

}
