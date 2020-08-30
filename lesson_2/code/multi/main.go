package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "Heal the world, make the better place")
	file.Close()
}