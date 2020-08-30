package main

import "os"

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}

	file.Write([]byte("any any monkey on the billding\n"))
	file.Write([]byte("any any monkey on the billding"))
	file.Close()
}