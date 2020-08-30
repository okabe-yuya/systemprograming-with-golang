package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "applicant/json")
	// to json
	source := map[string]string{
		"Hello": "world",
		"Feel": "Good",
	}
	// gzip writer
	g := gzip.NewWriter(w)
	writer := io.MultiWriter(g, os.Stdout)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	encoder.Encode(source)
	g.Close()
}



func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}