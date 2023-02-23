package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/araujo88/grpc-go-tutorial/primes/proto"
)

var addr string = "localhost:50051"

func computePrimeFactors(c pb.PrimesServiceClient) {
	log.Println("computePrimeFactors was invoked")

	req := &pb.PrimesRequest{
		Integer: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Factor: %d\n", msg.Prime)
	}
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewPrimesServiceClient(conn)

	//doGreet(c)

	computePrimeFactors(c)
}
