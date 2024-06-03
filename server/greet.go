package main

import (
	"context"
	"log"

	proto "github.com/neyaadeez/grpc/proto"
)

func (s *Server) Greet(conx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("Greet function was invoked with parameter %v\n", in)
	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
