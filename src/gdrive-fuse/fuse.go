package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
)

type FuseI interface {
	List() ([]File, error)
	Write(filename string, r io.Reader) error
	Read(filname string, w io.Writer) error
	Create(filename string) (string, error)
	Delete(filname string) error
}

// FileSystem :
type FileSystem struct {
	conn       *fuse.Conn
	mountpoint string
	fuse       FuseI
}

// Folder :
type Folder struct {
	Name string
	fuse FuseI
}

// File :
type File struct {
	Name string
	Id   string
	Size uint64
	fuse FuseI
}

// NewFileSystem :
func NewFileSystem(mountpoint string, f FuseI) (*FileSystem, error) {

	if mountpoint == "" {
		return nil, errors.New("mountpoint cannot be an empty string")
	}

	fmt.Printf("mounting at %s\n", mountpoint)

	c, err := fuse.Mount(
		mountpoint,
		// fuse.FSName("filesystem 1"),
		// fuse.Subtype("subtype 1"),
		// fuse.LocalVolume(),
		// fuse.VolumeName("Volume name"),
	)
	if err != nil {
		return nil, err
	}

	return &FileSystem{c, mountpoint, f}, nil
}

// Serve :
func Serve(f *FileSystem) {

	go func() {
		err := fs.Serve(f.conn, f)
		if err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	defer fuse.Unmount(f.mountpoint)

	<-stop
	fmt.Printf("\n%s unmounted\n", f.mountpoint)
}

// Root :
func (f *FileSystem) Root() (fs.Node, error) {
	fmt.Println("Root")
	return &Folder{"root", f.fuse}, nil
}

// Attr :
func (f *Folder) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0777
	return nil
}

// Lookup :
func (f *Folder) Lookup(ctx context.Context, name string) (fs.Node, error) {

	files, err := f.fuse.List()

	if err != nil {
		return nil, err
	}

	for _, file := range files {

		if file.Name == name {
			return &File{file.Name, file.Id, file.Size, f.fuse}, nil
		}
	}

	return nil, fuse.ENOENT
}

// ReadDirAll :
func (f *Folder) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {

	files, err := f.fuse.List()
	if err != nil {
		return nil, err
	}

	dirEntries := make([]fuse.Dirent, len(files))

	for i, file := range files {
		dirEntries[i] = fuse.Dirent{Inode: uint64(i), Name: file.Name, Type: fuse.DT_File}
	}

	return dirEntries, nil
}

//Create :
func (f *Folder) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
	log.Println("Create request for name", req.Name)

	id, err := f.fuse.Create(req.Name)
	if err != nil {
		return nil, nil, err
	}
	file := &File{req.Name, id, 0, f.fuse}

	return file, file, nil
}

func (f *Folder) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	log.Println("Remove request for ", req.Name)

	err := f.fuse.Delete(req.Name)
	if err != nil {
		return err
	}
	return nil
}

// Attr :
func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {

	//fmt.Printf("File Attr: size %d ID %s\n", f.Size, f.Id)
	a.Inode = 2
	a.Mode = 0777
	a.Size = f.Size

	return nil
}

// Read:
func (f *File) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	log.Println("Requested Read on File", f.Name)
	//fuseutil.HandleRead(req, resp, f.data)
	return nil
}

// ReadAll :
func (f *File) ReadAll(ctx context.Context) ([]byte, error) {

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	f.fuse.Read(f.Id, writer)
	return b.Bytes(), nil
}

// WriteAll :
func (f *File) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {

	fmt.Printf("** writing %d to %s\n", len(req.Data), f.Name)

	r := bytes.NewReader(req.Data)
	err := f.fuse.Write(f.Name, r)
	if err != nil {
		return err
	}
	resp.Size = len(req.Data)
	return nil
}

func (f *File) Flush(ctx context.Context, req *fuse.FlushRequest) error {
	//log.Println("Flushing file", f.Name)
	return nil
}
func (f *File) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	//log.Println("Open call on file", f.Name)
	return f, nil
}

func (f *File) Release(ctx context.Context, req *fuse.ReleaseRequest) error {
	//log.Println("Release requested on file", f.Name)
	return nil
}

func (f *File) Fsync(ctx context.Context, req *fuse.FsyncRequest) error {
	//log.Println("Fsync call on file", f.Name)
	return nil
}
