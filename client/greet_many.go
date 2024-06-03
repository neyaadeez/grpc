package main

import (
	"context"
	"io"
	"log"

	proto "github.com/neyaadeez/grpc/proto"
)

func doGreetMany(c proto.GreetServiceClient) {
	log.Printf("Do manyGreet was invoked")

	stream, err := c.GreetMany(context.Background(), &proto.GreetRequest{
		FirstName: "mustafa",
	})

	if err != nil {
		log.Fatalf("error while making request: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving data form server: %v\n", err)
		}

		log.Printf("Received data from server: %s\n", res.Result)

	}
}
