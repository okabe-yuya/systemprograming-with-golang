package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		// send channel to close
		time.Sleep(time.Second * 2)
		cancel()
	}()

	if v, ok := <- ctx.Done(); ok {
		fmt.Println("unsafed close")
	} else {
		fmt.Println("value: ", v)
	}
}