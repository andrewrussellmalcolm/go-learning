package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"streaming/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	address = flag.String("address", "localhost:10000", "The server address")
)

func main() {
	flag.Parse()

	fmt.Println(*address)

	creds, err := credentials.NewClientTLSFromFile("server.pem", "")
	if err != nil {
		log.Fatalf("cert load error: %s", err)
	}

	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	streamingServiceClient := streamingservice.NewStreamingServiceClient(conn)

	ctx := context.Background()

	client, err := streamingServiceClient.GetStream(ctx, &streamingservice.Void{})

	if err != nil {
		panic(err)
	}

	for {
		frame, err := client.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v %v", client, err)
		}

		fmt.Printf("Frame %d payload size=%d payload=%s\n", frame.Index, len(frame.Payload), string(frame.Payload))
	}

	fmt.Println("Frame transfer complete")
}
