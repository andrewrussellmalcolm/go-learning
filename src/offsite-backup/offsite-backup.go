package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

/** */
func main() {

	if len(os.Args) < 2 {
		fmt.Println("Must supply a file to upload")
		os.Exit(-1)
	}
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	srv, err := drive.New(getClient(config))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	upload(srv, os.Args[1])
	//list(srv)
}

/** */
func list(srv *drive.Service) {
	r, err := srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}

/** */
func upload(srv *drive.Service, filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening %q: %v", filename, err)
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("Could not obtain stat %v", err)
	}

	size := fi.Size()

	fmt.Printf("The file is %d bytes long\n", size)

	t0 := time.Now()

	r := NewProgressReader(f, func(chunk, total int64) {

		if chunk > 0 {
			fmt.Printf("Chunk of %d bytes transferred, total %d, time elapsed %.0fs\n", chunk, total, time.Since(t0).Seconds())
		}
	})

	driveFile, err := srv.Files.Create(&drive.File{Name: filename, MimeType: "application/tar"}).Media(r).Do()
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}

	_ = driveFile

	fmt.Printf("Total bytes transferred %d\n", r.total)
	t1 := time.Now()

	time := t1.Sub(t0).Seconds()
	fmt.Printf("time taken=%.1fs\n", time)

	fmt.Printf("transfer rate %.2fMbytes/s\n", float32(size)/float32(time)/1000000.0)
}
