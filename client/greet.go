package main

import (
	"context"
	"log"

	proto "github.com/neyaadeez/grpc/proto"
)

func doGreet(c proto.GreetServiceClient) {
	log.Printf("inVoked greet service client")
	res, err := c.Greet(context.Background(), &proto.GreetRequest{
		FirstName: "mustafa",
	})

	if err != nil {
		log.Fatalf("error while making request to the server: %v", err)
	}

	log.Printf("Response from the server: %s", res.Result)
}
