package main

import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	// os.Stdout.Write([]byte(buffer.buf))
	buffer.Flush()
	buffer.WriteString("example\n")
	// os.Stdout.Write([]byte(buffer.String()))
	buffer.Flush()
}