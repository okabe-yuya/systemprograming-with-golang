package main

import (
	"encoding/binary"
	"hash/crc32"
	"bytes"
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
	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	chunks := readChunks(file)
	// write Signature
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// write IDHR
	io.Copy(newFile, chunks[0])
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
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

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	// calcuate crc
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk %v, (%d bytes) \n", string(buffer), length)

	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}