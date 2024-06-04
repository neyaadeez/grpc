package main

import (
	"fmt"
	"io"
	"log"

	proto "github.com/neyaadeez/grpc/proto"
)

func (s *Server) GreetEveryOne(stream proto.GreetService_GreetEveryOneServer) error {
	log.Println("invoked server longGreet (client streaming)")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while receiving data: %v\n", err)
		}

		res = fmt.Sprintf("Hello %s!\n", req.FirstName)

		err = stream.Send(&proto.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("error while sending: %v\n", err)
		}
	}

}
