package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// insert package file >> file.go
	io.WriteString(file, "package file\n")
	io.Copy(os.Stdout, file)
}