package main

import (
	pb "github.com/zaenalarifin12/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:5000"

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	if c == nil {
		log.Fatalf("Failed to create client")
	}
	//doGreet(c)
	doGreetManyTimes(c)
}
