package main

import (
	"io"
	"os"
	"fmt"
	"strings"
	"bytes"
)

func main() {
	w := os.Stdout
	r := strings.NewReader("ABCDE")
	CopyN(w, r, 2)
}

func CopyN(des io.Writer, src io.Reader, length int) (int64, error) {
	buf := make([]byte, length)
	src.Read(buf)

	r := bytes.NewReader(buf)
	w, err := io.Copy(des, r)
	if err != nil {
		return 0, err
	}
	return w, nil
}

