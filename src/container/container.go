package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {

	var cmd *exec.Cmd

	if len(os.Args) > 1 {
		cmd = exec.Command(os.Args[1], strings.Join(os.Args[2:], " "))

	} else {
		cmd = exec.Command("bash")
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := syscall.Chroot("./jail")
	if err != nil {
		fmt.Printf("Failed to chroot %v\n", err)
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
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
