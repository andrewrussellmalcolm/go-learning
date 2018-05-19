package main

import (
	"cache/api"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	address = flag.String("address", "localhost:10000", "The server address")
)

func main() {
	flag.Parse()

	fmt.Printf("contacting %s\n", *address)

	conn, err := grpc.Dial(*address, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	cacheServiceClient := cacheservice.NewCacheServiceClient(conn)

	putMeasurements(cacheServiceClient)

	//getMeasurements(cacheServiceClient)

}

func putMeasurements(cacheServiceClient cacheservice.CacheServiceClient) {

	ctx := context.Background()
	c, err := cacheServiceClient.PutMeasurementStream(ctx)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		err = c.Send(&cacheservice.Measurement{0, 0, time.Now().UnixNano()})
		if err != nil {
			panic(err)
		}
	}

	c.CloseSend()

	time.Sleep(time.Second)
}

func getMeasurements(cacheServiceClient cacheservice.CacheServiceClient) {
	gmr := cacheservice.GetMeasurementRequest{}
	ctx := context.Background()
	s, err := cacheServiceClient.GetMeasurementStream(ctx, &gmr)
	if err != nil {
		panic(err)
	}

	for {
		m, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v %v", s, err)
		}

		fmt.Printf("Measurement ID=%d  value=%d= timestamp=%d\n", m.GetSensorLocationID(), m.GetValue(), m.GetTimestamp())
	}

	fmt.Println("End of measurement stream")
}
