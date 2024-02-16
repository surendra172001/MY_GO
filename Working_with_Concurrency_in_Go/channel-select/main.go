package main

import (
	"fmt"
	"time"
)

func server1(chan1 chan<- string) {
	for {
		time.Sleep(6 * time.Second)
		chan1 <- "Server1 message"
	}
}

func server2(chan2 chan<- string) {
	for {
		time.Sleep(3 * time.Second)
		chan2 <- "Server2 message"
	}
}

// print welcome message
// create 2 channels one for receiving from server1 and other from server2
// create 2 functions and then corresponding 2 routines
// loop infinitely and take data from server1, server2
// write a select statement inside the loop and have 4 cases, 2 cases for each server
// close the channels
func main() {
	fmt.Println("Select statement demonstration")
	fmt.Println("------------------------------")

	chan1 := make(chan string)
	chan2 := make(chan string)

	go server1(chan1)
	go server2(chan2)

	for {
		select {
		case c1 := <-chan1:
			fmt.Printf("case 1: %s\n", c1)
		case c2 := <-chan1:
			fmt.Printf("case 2: %s\n", c2)
		case c3 := <-chan2:
			fmt.Printf("case 3: %s\n", c3)
		case c4 := <-chan2:
			fmt.Printf("case 4: %s\n", c4)
		}
	}

	// fmt.Println("End of the channel - select demo")

}
