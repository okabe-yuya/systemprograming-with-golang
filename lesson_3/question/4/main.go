package main

import (
	"net/http"
	"archive/zip"
)

func main() {
	http.HandleFunc("/zip", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=download.zip")
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	var files = []struct {
		Name string
		Body string
	}{
		{"README.md", "Hello Every one"},
		{"Install.md", "> brew upgrade"},
		{"Usage.md", "Chek official web page"},
	}

	for _, file := range files {
		f, err := zipWriter.Create(file.Name)
		if err != nil {
			panic(err)
		}
		f.Write([]byte(file.Body))
	}
}