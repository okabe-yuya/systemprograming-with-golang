package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	// retry and loop
	for {
		var err error
		if conn == nil {
			// dial to server
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// send message on POST
		request, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		if err := request.Write(conn); err != nil {
			panic(err)
		}
		// read response from server
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}