package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("Working with simple channels")
	fmt.Println("----------------------------")

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type someting and press ENTER(enter Q to quit)")

	for {
		fmt.Print("--> ")
		var inputStr string
		_, err := fmt.Scanln(&inputStr)

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(inputStr) != "q" {
			ping <- inputStr
			fmt.Println("Pong says", <-pong)
		} else {
			break
		}
	}

	fmt.Println("Ending of simple channels demo")
	fmt.Println("------------------------------")

	close(ping)
	close(pong)
}

func shout(ping <-chan string, pong chan<- string) {
	for {
		if s, ok := <-ping; ok {
			pong <- fmt.Sprintf("%s !!!", strings.ToUpper(s))
		}
	}
}
