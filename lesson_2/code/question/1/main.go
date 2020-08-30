package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}

	file.Write([]byte(fmt.Sprintf("You wirte %v", time.Now())))
	file.Write([]byte(fmt.Sprintf("You wirte %v", 114514)))
	file.Close()
}
