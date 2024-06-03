package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/neyaadeez/grpc/proto"
)

func (s *Server) Greet(conx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("Greet function was invoked with parameter %v\n", in)
	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) GreetMany(in *proto.GreetRequest, stream proto.GreetService_GreetManyServer) error {
	log.Printf("GreetMany function was invoked with parameter: %v\n", in)

	for i := 0; i < 10; i++ {
		stream.Send(&proto.GreetResponse{
			Result: fmt.Sprintf("Hello %s, this is steaming test: %d\n", in.FirstName, i),
		})
	}

	return nil
}
