package main

import (
	"io"
	"os"
	"net"
	"net/http"
	"bufio"
	"fmt"
)

// not parse
// func main() {
// 	conn, err := net.Dial("tcp", "google.com:80")
// 	if err != nil {
// 		panic(err)
// 	}
// 	conn.Write([]byte("GET / HTTP/1.0\r\nHost: google.com\r\n\r\n"))
// 	io.Copy(os.Stdout, conn)
// }

// parse
func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Header)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}