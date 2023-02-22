package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/araujo88/grpc-go-tutorial/sum/proto"
)

var addr string = "localhost:50051"

func getSum(c pb.SumServiceClient) {
	log.Println("getSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 1,
		Num2: 2,
	})

	if err != nil {
		log.Fatalf("Could not obtain result: %v\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewSumServiceClient(conn)

	getSum(c)
}
