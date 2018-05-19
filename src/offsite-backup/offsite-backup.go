package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

const description = `
=======================================================================
backup a single file to google drive

credentials come from an external file, obtained from here:
https://console.developers.google.com/flows/enableapi?apiid=drive
You must download the resulting file and save it in client-secret.json 
in the same directory as the executable
The first time it is run, you will be prompted to follow a link to 
obtain a token, which must be save to a file called token.json, 
again in the same directory as the executable
=======================================================================
`

var (
	flagFile = flag.String("upload", "", "The file to upload")
	flagList = flag.Bool("list", false, "List uploaded files")
	flagTime = flag.Int("timeout", 5, "Permitted time for upload operation (s)")
)

/** */
func main() {

	flag.Parse()

	b, err := ioutil.ReadFile("client-secret.json")
	if err != nil {
		fmt.Println("Unable to read client secret file")
		fmt.Println(description)
		os.Exit(-1)
	}

	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		fmt.Println("Unable to parse client secret file to config")
		fmt.Println(description)
		os.Exit(-1)
	}

	srv, err := drive.New(getClient(config))
	if err != nil {
		fmt.Println("Unable to retrieve Drive client")
		fmt.Println(description)
		os.Exit(-1)
	}

	if *flagList {
		list(srv)
	} else if *flagFile != "" {
		upload(srv, *flagFile)
	} else {
		fmt.Println("No arguments supplied. Usage is:")
		flag.PrintDefaults()
		fmt.Println(description)
	}
}

/** */
func list(srv *drive.Service) {
	r, err := srv.Files.List().Q("trashed=false").Do()
	if err != nil {
		log.Fatalf("Unable to list google drive files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found")
	} else {
		for _, i := range r.Files {
			fmt.Printf("Filename %s ID (%s)\n", i.Name, i.Id)
		}
	}
}

/** */
func createBackupFolder(srv *drive.Service) (*drive.File, error) {
	r, err := srv.Files.List().Q("trashed=false").Do()
	if err != nil {
		log.Fatalf("Unable to list google drive files: %v", err)
	}

	for _, f := range r.Files {
		if f.Name == "Backup" {
			return f, nil
		}
	}

	driveFolder, err := srv.Files.Create(&drive.File{Name: "Backup", MimeType: "application/vnd.google-apps.folder"}).Do()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Creating Backup folder with id %s\n", driveFolder.Id)
	return driveFolder, nil
}

/** */
func upload(srv *drive.Service, filename string) {

	driveFolder, err := createBackupFolder(srv)

	if err != nil {
		log.Fatalf("Unable to create folder: %v", err)
	}

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

	totalTransferred := int64(0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*flagTime))
	defer cancel()

	uploadName := time.Now().Format("2006-01-02[15:04:05]") + "." + filename
	t0 := time.Now()
	driveFile, err := srv.Files.Create(&drive.File{Name: uploadName, MimeType: "application/tar", Parents: []string{driveFolder.Id}}).
		Media(f).
		Context(ctx).
		Fields().
		ProgressUpdater(func(current, total int64) {

			fmt.Printf("%d bytes transferred, time elapsed %.0fs\n", current, time.Since(t0).Seconds())

			totalTransferred = current
		}).
		Do()

	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}

	_ = driveFile
	t1 := time.Now()

	time := t1.Sub(t0).Seconds()

	fmt.Printf("Total bytes transferred %d\n", totalTransferred)
	fmt.Printf("Time taken=%.1fs\n", time)
	fmt.Printf("Transfer rate %.2fMbytes/s\n", float32(size)/float32(time)/1000000.0)
	fmt.Printf("File %s uploaded as %s\n", filename, uploadName)
}
