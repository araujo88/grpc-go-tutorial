package main

import (
	"log"
	"net"

	pb "github.com/araujo88/grpc-go-tutorial/primes/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.PrimesServiceServer
}

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.PrimesService_PrimesServer) error {
	log.Printf("Primes was invoked with %v\n", in)

	var k uint32 = 2
	N := in.Integer

	for i := N; i > 1; i-- {

		if N%k == 0 {
			N = N / k
			stream.Send(&pb.PrimesResponse{
				Prime: k,
			})
		} else {
			k += 1
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterPrimesServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
