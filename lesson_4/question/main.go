package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.After(3 * time.Second)
	<- ch
	fmt.Println("Elapsed time!!")
}
