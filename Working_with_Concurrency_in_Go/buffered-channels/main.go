package main

import (
	"fmt"
	"time"
)

func listenToChan(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Buffered channel demonstration")
	fmt.Println("------------------------------")

	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 20; i++ {
		fmt.Println("Sending", i, "to channel...")
		ch <- i
		fmt.Println("Sent", i, "to channel...")
	}
	close(ch)
}
