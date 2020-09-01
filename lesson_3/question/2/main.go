package main

import (
	"crypto/rand"
	"bufio"
	"os"
	"encoding/hex"
)

func main() {
	file, err := os.Create("random.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := bufio.NewWriter(file)
	maxByte := 1024
	total := 0
	for (total <= maxByte) {
		a := make([]byte, 64)
		rand.Read(a)
		buffer.WriteString(hex.EncodeToString(a))
		buffer.WriteString("\n")
		buffer.Flush()
		total += 64
	}
}