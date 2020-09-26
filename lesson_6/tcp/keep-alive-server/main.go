package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"io"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// after accepted, loop receive request
			for {
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				// read request
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					netErr, ok := err.(net.Error)
					if ok && netErr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				content := "Hello World\n"

				response := http.Response{
					StatusCode: 200,
					ProtoMajor: 1,
					ProtoMinor: 1,
					ContentLength: int64(len(content)),
					Body: ioutil.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}