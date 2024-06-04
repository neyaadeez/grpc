package main

import (
	"fmt"
	"io"
	"log"

	proto "github.com/neyaadeez/grpc/proto"
)

func (s *Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	log.Println("invoked server longGreet (client streaming)")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("error while receiving data: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}

}
