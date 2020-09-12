package main

import (
	"strings"
	"io"
	"os"
	"bytes"
)

var (
	computer = strings.NewReader("COMPUTER")
	system = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader
	buffer := make([]byte, 5)
	buf := make([]byte, 1)
	// A
	if _, err := programming.ReadAt(buf, 5); err != nil {
		panic(err)
	}
	buffer[0] = buf[0]
	// S
	if _, err := system.ReadAt(buf, 0); err != nil {
		panic(err)
	}
	buffer[1] = buf[0]
	// C
	if _, err := computer.ReadAt(buf, 0); err != nil {
		panic(err)
	}
	buffer[2] = buf[0]
	// I
	if _, err := programming.ReadAt(buf, 8); err != nil {
		panic(err)
	}
	buffer[3] = buf[0]
	buffer[4] = buf[0]

	stream = bytes.NewReader(buffer)
	io.Copy(os.Stdout, stream)
}