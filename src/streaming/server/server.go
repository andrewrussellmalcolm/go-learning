package main

import (

	//"errors"

	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"streaming/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type streamingService struct {
	frameIndex int32
}

var (
	port   = flag.Int("port", 10000, "The server port")
	noAuth = flag.Bool("noauth", false, "Set to true to disable authentication")
	frames = flag.Int("frames", 20, "The number of frames returned in a single request")
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("server.pem", "server.key")
	if err != nil {
		log.Fatalf("Failed to setup tls: %v", err)
	}

	var grpcServer *grpc.Server

	if *noAuth {
		grpcServer = grpc.NewServer()
	} else {
		grpcServer = grpc.NewServer(grpc.Creds(creds))
	}

	streamingservice.RegisterStreamingServiceServer(grpcServer, &streamingService{})

	fmt.Printf("server listening on %d\n", *port)
	fmt.Printf("frame per stream %d\n", *frames)
	fmt.Printf("no auth %t\n", *noAuth)
	grpcServer.Serve(lis)
}

func (t *streamingService) GetStream(void *streamingservice.Void, stream streamingservice.StreamingService_GetStreamServer) error {

	payload := make([]byte, 60)

	for t.frameIndex < int32(*frames) {

		for i := 0; i < len(payload); i++ {

			payload[i] = byte('A') + byte(rand.Intn(26))
		}

		frame := streamingservice.Frame{Index: t.frameIndex, Payload: payload}

		if err := stream.Send(&frame); err != nil {
			return err
		}
		fmt.Printf("frame %d payload size = %d payload=%s\n", frame.Index, len(frame.Payload), string(frame.Payload))
		t.frameIndex++
	}
	t.frameIndex = 0

	fmt.Println("Frame transfer complete")
	return nil
}
