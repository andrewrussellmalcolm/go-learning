package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"streaming/api"

	"github.com/gotk3/gotk3/gdk"

	"github.com/gotk3/gotk3/gtk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//go build -v -tags gtk_3_18 -gcflags "-N -l"
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

	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Webcam Client")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	paneView, err := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
	label, err := gtk.LabelNew("$$")
	paneView.Add(label)

	image, err := gtk.ImageNew()
	if err != nil {
		panic(err)
	}

	win.Add(image)
	win.SetDefaultSize(800, 200)
	win.ShowAll()
	win.GetFocus()

	payload := make(chan []byte)

	go streamImages(conn, payload)

	exit := false
	for exit == false {
		select {
		case payload := <-payload:
			updateImageFromFramePayload(image, payload)
		default:
		}

		if !win.IsVisible() {
			exit = true
		} else {
			gtk.MainIterationDo(false)
		}
	}
}

func streamImages(conn *grpc.ClientConn, payload chan []byte) {

	streamingServiceClient := streamingservice.NewStreamingServiceClient(conn)

	ctx := context.Background()

	client, err := streamingServiceClient.GetStream(ctx, &streamingservice.Void{})

	if err != nil {
		log.Fatalf("%v %v", client, err)
	}

	for {
		frame, err := client.Recv()
		if err == io.EOF {
			close(payload)
			break
		}
		if err != nil {
			log.Fatalf("%v %v", client, err)
		}

		payload <- frame.Payload
	}
}

func updateImageFromFramePayload(image *gtk.Image, data []byte) {

	x := gdk.PixbufGetFormats()

	for _, y := range x {
		d, _ := y.GetDescription()
		fmt.Printf("%v\n", d)
	}

	loader, err := gdk.PixbufLoaderNew()
	if err != nil {
		panic(err)
	}

	_, err = loader.Write(data)
	if err != nil {
		panic(err)
	}

	pixbuf, err := loader.GetPixbuf()
	if err != nil {
		panic(err)
	}
	image.SetFromPixbuf(pixbuf)

	loader.Close()
}
