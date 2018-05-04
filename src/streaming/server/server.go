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
	port = flag.Int("port", 10000, "The server port")
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
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	streamingservice.RegisterStreamingServiceServer(grpcServer, &streamingService{})

	grpcServer.Serve(lis)
}

func (t *streamingService) GetStream(void *streamingservice.Void, stream streamingservice.StreamingService_GetStreamServer) error {

	payload := make([]byte, 60)

	for t.frameIndex < 10 {

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
