package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"streaming/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	address = flag.String("address", "localhost:10000", "The server address")
	noAuth  = flag.Bool("noauth", false, "Set to true to disable authentication")
)

func main() {
	flag.Parse()

	fmt.Printf("contacting %s\n", *address)
	fmt.Printf("no auth %t\n", *noAuth)

	creds, err := credentials.NewClientTLSFromFile("server.pem", "")
	if err != nil {
		log.Fatalf("cert load error: %s", err)
	}

	var conn *grpc.ClientConn

	if *noAuth {
		conn, err = grpc.Dial(*address, grpc.WithInsecure())
	} else {
		conn, err = grpc.Dial(*address, grpc.WithTransportCredentials(creds))
	}
	if err != nil {
		log.Fatalf("dial failed %v", err)
	}

	streamingServiceClient := streamingservice.NewStreamingServiceClient(conn)

	ctx := context.Background()

	client, err := streamingServiceClient.GetStream(ctx, &streamingservice.Void{})

	if err != nil {
		log.Fatalf("%v %v", client, err)
	}

	f, err := os.OpenFile("tmp.jpg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)

	if err != nil {
		log.Fatalf("file open failed %v", err)
	}

	for {
		frame, err := client.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v %v", client, err)
		}

		fmt.Printf("Frame %d payload size=%d\n", frame.Index, len(frame.Payload))

		_, err = f.Write(frame.Payload)
		if err != nil {
			log.Fatalf("file write failed %v", err)
		}
	}

	f.Close()

	fmt.Println("Frame transfer complete")
}
