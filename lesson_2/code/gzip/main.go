package main

import (
	"compress/gzip"
	"os"
	"io"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example \n")
	writer.Close()
	file.Close()
}