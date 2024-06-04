package main

import (
	"context"
	"log"
	"time"

	proto "github.com/neyaadeez/grpc/proto"
)

func doLongGreet(c proto.GreetServiceClient) {
	log.Printf("client streaming function is invoked\n")

	requests := []*proto.GreetRequest{
		{FirstName: "mustafa"},
		{FirstName: "rahul"},
		{FirstName: "rohit"},
		{FirstName: "john"},
		{FirstName: "downy"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("unable to send request to the server: %v\n", err)
	}

	for _, req := range requests {
		log.Printf("sending stream of request to the server: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from the server: %v\n", err)
	}

	log.Printf("Response from the server: %s\n", res.Result)
}
