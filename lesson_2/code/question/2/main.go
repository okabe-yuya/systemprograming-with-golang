package main

import (
	"encoding/csv"
	"os"
	// "io"
)

func main() {
	file, err := os.Create("sample.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	val := []string{"sample", "\n"}
	writer.Write(val)
	writer.Flush()
	writer.Write(val)
	writer.Flush()
}