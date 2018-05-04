package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {

	var path = "jail"
	var prog string

	if len(os.Args) > 1 {
		prog = os.Args[1]
	} else {
		fmt.Println("Must supply a progam to put in jail")
		os.Exit(0)
	}

	cmd := exec.Command("./"+prog, strings.Join(os.Args[2:], " "))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	syscall.Rmdir(path)
	err := syscall.Mkdir(path, 0777)
	if err != nil {
		fmt.Printf("Failed to mkdir %s %v\n", path, err)
		os.Exit(10)
	}

	err = cp(path+"/"+prog, prog)
	if err != nil {
		fmt.Printf("Failed to copy %s %v\n", prog, err)
		os.Exit(10)
	}

	err = os.Chmod(path+"/"+prog, 0500)
	if err != nil {
		fmt.Printf("Failed to chmod %s\n", err)
		os.Exit(10)
	}

	err = syscall.Chroot(path)
	if err != nil {
		fmt.Printf("Failed to chroot %s %v\n", path, err)
		os.Exit(10)
	}

	err = os.Chdir("/")
	if err != nil {
		fmt.Printf("Failed to chdir %v\n", err)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start %v\n", err)
		os.Exit(10)
	}

	fmt.Printf("Started and waiting for exit\n")
	err = cmd.Wait()
	if err == nil {
		fmt.Printf("Command completed with no error\n")
	} else {
		fmt.Printf("Command completed: %s\n", err)
	}

	err = syscall.Unlink(prog)
	if err != nil {
		fmt.Printf("Failed to unlink %s %v\n", path, err)
		os.Exit(10)
	}
}

func cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}
