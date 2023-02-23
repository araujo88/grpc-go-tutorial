package main

import (
	"context"
	"log"
	"time"

	pb "github.com/araujo88/grpc-go-tutorial/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Leonardo"},
		{FirstName: "Eduardo"},
		{FirstName: "Test"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
