package main

import (
	"log"

	proto "github.com/neyaadeez/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to connect to the server, error: %v", err)
	}

	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)
	// doGreet(client)
	// doGreetMany(client)
	// doLongGreet(client)
	doGreetEveryone(client)

}
