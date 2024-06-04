package main

import (
	"context"
	"io"
	"log"
	"time"

	proto "github.com/neyaadeez/grpc/proto"
)

func doGreetEveryone(c proto.GreetServiceClient) {
	log.Println("bidirection client request is invoked")

	requests := []*proto.GreetRequest{
		{FirstName: "mustafa"},
		{FirstName: "rahul"},
		{FirstName: "rohit"},
		{FirstName: "john"},
		{FirstName: "downy"},
	}

	stream, err := c.GreetEveryOne(context.Background())

	if err != nil {
		log.Fatalf("unable to send request to the server: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("sending stream of request to the server: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
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
		close(waitc)
	}()

	<-waitc
}
