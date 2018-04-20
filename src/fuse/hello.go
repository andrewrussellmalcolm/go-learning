package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
)

// Main :
func main() {

	var mountpoint string

	flag.StringVar(&mountpoint, "mp", "", "defines the mount point for this filesystem")
	flag.Parse()

	if mountpoint == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	fmt.Printf("mounting at %s\n", mountpoint)

	c, err := fuse.Mount(
		mountpoint,
		fuse.FSName("filesystem 1"),
		fuse.Subtype("subtype 1"),
		fuse.LocalVolume(),
		fuse.VolumeName("Volume name"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go func() {
		err = fs.Serve(c, FS{})
		if err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	defer fuse.Unmount(mountpoint)

	<-stop
	fmt.Printf("\n%s unmounted\n", mountpoint)
}

// FS implements the file system.
type FS struct{}

// Root :
func (FS) Root() (fs.Node, error) {
	return Dir{}, nil
}

// Dir implements both Node and Handle for the root directory.
type Dir struct{}

// Attr :
func (Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0555
	return nil
}

// Lookup :
func (Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {

	for file := range files {

		if file.Name == name {
			return file, nil
		}
	}
	return nil, fuse.ENOENT
}

// File represents a file
type File struct {
	Name string
}

var files = map[File][]byte{
	File{"hello 1.txt"}: make([]byte, 0),
	File{"hello 2.txt"}: make([]byte, 0),
}

// ReadDirAll :
func (Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {

	dirEntries := make([]fuse.Dirent, len(files))

	i := 0
	for file := range files {
		dirEntries[i] = fuse.Dirent{Inode: uint64(i), Name: file.Name, Type: fuse.DT_File}
		i++
	}

	return dirEntries, nil
}

// Attr :
func (f File) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 2
	a.Mode = 0444
	a.Size = uint64(len(files[f]))
	return nil
}

// ReadAll :
func (f File) ReadAll(ctx context.Context) ([]byte, error) {
	return files[f], nil
}

// WriteAll :
func (f File) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) {

}
