package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func readChunks(file *os.File) []io.Reader {
	// list of chunks
	var chunks []io.Reader
	// skip to top 8byte
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		// 8byte:[signature], [4byte:[length], 4byte:[kind], Nbyte:[data], 4byte:[CRC]] ...
		if err := binary.Read(file, binary.BigEndian, &length); err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk %v, (%d bytes) \n", string(buffer), length)
}