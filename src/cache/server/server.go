package main

import (
	"encoding/binary"
	"time"

	//"errors"

	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"cache/api"

	"google.golang.org/grpc"
)

type cacheService struct {
	file *os.File
}

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {

	cs := cacheService{}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cs.file, err = os.OpenFile("/dev/shm/bf-cache", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed to open cache memory file %v", err)
	}

	//cs.file.Seek(1024*1024*1024*16, 0)

	fmt.Println("Starting")
	timeOperation(func() {
		for j := 0; j < 1024; j++ {
			for i := 0; i < 1024*1024; i++ {
				writeMeasurement(uint64(i+i*j), 0x1111222233334444, 0x5555666677778888, cs.file)
				fmt.Print(".")
			}
		}
	}, "")

	v0, t0, err := readMeasurement(0, cs.file)
	fmt.Printf("0x%x 0x%x\n", v0, t0)
	v1, t1, err := readMeasurement(1, cs.file)
	fmt.Printf("0x%x 0x%x\n", v1, t1)

	grpcServer := grpc.NewServer()

	cacheservice.RegisterCacheServiceServer(grpcServer, &cs)

	fmt.Printf("server listening on %d\n", *port)

	grpcServer.Serve(lis)
}

func (c *cacheService) GetMeasurement(ctx context.Context, in *cacheservice.GetMeasurementRequest) (*cacheservice.GetMeasurementResponse, error) {
	return nil, nil

}

// possible errors : cache full (or fail silently?)
func (c *cacheService) PutMeasurement(ctx context.Context, in *cacheservice.PutMeasurementRequest) (*cacheservice.Void, error) {
	return nil, nil

}

// for bulk upload at init time
func (c *cacheService) PutMeasurementStream(s cacheservice.CacheService_PutMeasurementStreamServer) error {

	fmt.Println("PMS")
	for {
		m, err := s.Recv()
		if err != nil {
			return err
		}
		fmt.Printf("Measurement ID=%d  value=%d= timestamp=%d\n", m.GetSensorLocationID(), m.GetValue(), m.GetTimestamp())
	}

	fmt.Println("PMS-E")
	return nil

}

// alternative for queries returning repeated data
func (c *cacheService) GetMeasurementStream(r *cacheservice.GetMeasurementRequest, s cacheservice.CacheService_GetMeasurementStreamServer) error {

	fmt.Println("GMS")
	m := cacheservice.Measurement{0, 0, 0}

	if err := s.Send(&m); err != nil {
		return err
	}

	fmt.Println("GMS E")

	return nil
}

//writeMeasurement(0, 0x1111222233334444, 0x5555666677778888, cs.file)
//writeMeasurement(1, 0x4444333322221111, 0x8888777766665555, cs.file)

func writeMeasurement(id, value, timestamp uint64, file *os.File) error {

	v := make([]byte, 8)
	t := make([]byte, 8)
	binary.BigEndian.PutUint64(v, value)
	binary.BigEndian.PutUint64(t, timestamp)

	vOff := int64(id * 16)
	tOff := int64(id*16 + 8)

	// use ID as offset into file, 16 bytes per record
	file.Seek(vOff, 0)
	_, err := file.Write(v)
	if err != nil {
		return err
	}
	file.Seek(tOff, 0)
	_, err = file.Write(t)
	if err != nil {
		return err
	}

	return nil
}

func readMeasurement(id uint64, file *os.File) (value, tiemstamp uint64, err error) {

	v := make([]byte, 8)
	t := make([]byte, 8)
	vOff := int64(id * 16)
	tOff := int64(id*16 + 8)

	// use ID as offset into file, 16 bytes per record
	file.Seek(vOff, 0)
	_, err = file.Read(v)
	if err != nil {
		return 0, 0, err
	}

	file.Seek(tOff, 0)
	_, err = file.Read(t)
	if err != nil {
		return 0, 0, err
	}

	return binary.BigEndian.Uint64(v), binary.BigEndian.Uint64(t), nil
}

func timeOperation(operation func(), statement string) {

	t0 := time.Now()
	operation()
	t1 := time.Now()

	time := float32(t1.Sub(t0).Nanoseconds())

	if time < 1000 {
		fmt.Printf("[%s] took %.1fns\n", statement, time)
	} else if time < 1000000 {
		fmt.Printf("[%s] took %.1fus\n", statement, time/1000.0)
	} else if time < 1000000000 {
		fmt.Printf("[%s] took %.1fms\n", statement, time/1000000.0)
	} else if time < 1000000000000 {
		fmt.Printf("[%s] took %.1fs\n", statement, time/1000000000.0)
	}
}
