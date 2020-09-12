package main

import (
	"os"
	"archive/zip"
	// "strings"
)

func main() {
	file, err := os.Create("newFile.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := zip.NewWriter(file)
	defer w.Close()

	var files = []struct {
		Name string
		Body string
	}{
		{"README.md", "Hello Every one"},
		{"Install.md", "> brew upgrade"},
		{"Usage.md", "Chek official web page"},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			panic(err)
		}
		// sr := strings.NewReader(file.Body)
		f.Write([]byte(file.Body))
	}
}
