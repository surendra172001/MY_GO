package main

import (
	"fmt"
	"sync"
)

var msg = "Hello, world!"
var wg sync.WaitGroup

func updateMessage(aMsg string) {
	defer wg.Done()
	msg = aMsg
}

func main() {
	wg.Add(2)
	go updateMessage("Hello, universe")
	go updateMessage("Hello, cosmos")
	wg.Wait()
	fmt.Println(msg)
}
