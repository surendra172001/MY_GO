package main

import (
	"fmt"
	"sync"
)

func printSomething(msg string, wg *sync.WaitGroup) {
	fmt.Println(msg)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	letters := []string{
		"alpha",
		"beta",
		"delta",
		"epsilon",
		"eta",
		"gamma",
		"pi",
		"theta",
		"zeta",
	}

	wg.Add(len(letters)) // adding count for every routine corresponding to every element in slice

	for i, str := range letters {
		go printSomething(fmt.Sprintf("This is letter: %v, %v", i, str), &wg)
	}

	wg.Wait() // waiting for all the go routines to complete

	fmt.Println("Go main function got completed")

}
