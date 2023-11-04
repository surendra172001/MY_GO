package main

import (
	"fmt"
	"sync"
	"time"
)

var buffSize = 5

var waitChans chan int

var customerNum = 0

var hairCutDuration = 500 * time.Millisecond
var reqGenInterval = 250 * time.Millisecond
var customerCnt = 15

var wg sync.WaitGroup

func generateRequests() {
	for i := 0; i < customerCnt; i++ {
		fmt.Printf("Customer %d has arrived\n", i)
		if len(waitChans) == cap(waitChans) {
			fmt.Printf("There is no place in the waiting room customer %d is leaving\n", i)
		} else {
			waitChans <- customerNum
			customerNum++
		}
		fmt.Println("len waitchans ", len(waitChans))
		time.Sleep(reqGenInterval)
	}
	waitChans <- -1
	wg.Done()
}

func processRequests() {
	for i := range waitChans {
		// fmt.Printf("customer num: %d\n", i)
		if i == -1 {
			fmt.Print("Closing the shop\n")
			break
		}
		fmt.Printf("Doing haircut for customer %d\n", i)
		time.Sleep(hairCutDuration)
		fmt.Printf("Haircut for customer %d is done\n", i)
	}
	wg.Done()
}

func main() {
	waitChans = make(chan int, buffSize)

	wg.Add(2)

	go generateRequests()

	go processRequests()

	wg.Wait()

	fmt.Printf("The channel size is %d\n", len(waitChans))
	close(waitChans)
}
