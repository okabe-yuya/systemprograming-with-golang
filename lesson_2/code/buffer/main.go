package main

import (
	"bytes"
	// "fmt"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	buffer.Write([]byte("bytes.Buffer apple\n"))
	buffer.Write([]byte("bytes.Buffer banana\n"))
	buffer.Write([]byte("bytes.Buffer orange\n"))
	buffer.Write([]byte("bytes.Buffer melon\n"))
	// fmt.Println(buffer.String())

	// Stdoutでもbufferingしたデータを出力することが可能
	os.Stdout.Write([]byte(buffer.String()))
}