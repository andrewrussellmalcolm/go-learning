package main

import (

	//"errors"

	"flag"
	"fmt"
	"log"
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

	initWebcam()
	defer closeWebcam()
	startStreaming()

	fmt.Println("A")
	for t.frameIndex < int32(*frames) {

		frameMessage := streamingservice.Frame{Index: t.frameIndex}
		waitForFrame()

		payload, err := getFrame()

		_ = err
		frameMessage.Payload = payload

		if err := stream.Send(&frameMessage); err != nil {
			return err
		}

		fmt.Printf("frame %d payload size = %d payload=%s\n", frameMessage.Index, len(frameMessage.Payload))
		t.frameIndex++
	}
	t.frameIndex = 0

	fmt.Println("Frame transfer complete")
	return nil
}
