package main

import (
	"log"
	"net"

	proto "github.com/neyaadeez/grpc/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("problem with server, failed to listen, error: %v\n", err)
	}

	log.Printf("Listening on port: %v", addr)

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("problem with setting up grpc server, error: %v\n", err)
	}

}
