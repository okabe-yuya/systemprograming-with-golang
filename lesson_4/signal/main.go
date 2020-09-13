package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := create()
	signal.Notify(signals, syscall.SIGINT)

	// wait singal
	fmt.Println("Waiting SIGINT (CTRL + C")
	<- signals
	fmt.Println("SIGINT arrived")
}

func create(buffers ...int) (chan os.Signal) {
	var ch chan os.Signal
	switch len(buffers) {
	case 0:
		ch = _create(1)
	case 1:
		ch = _create(buffers[0])
	}
	return ch
}

func _create(buf int) (chan os.Signal) {
	return make(chan os.Signal, buf)
}