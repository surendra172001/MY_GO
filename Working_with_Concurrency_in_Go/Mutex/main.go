package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup

func main() {
	var mtx sync.Mutex
	// create a variable for balance tracking
	bankBalance := 0

	// create weekly income sources
	sources := []Income{
		{"Main job", 500},
		{"Pocket Money", 10},
		{"Part time job", 100},
		{"Dog walking", 50},
	}

	// predict balance across 52 weeks for each income source

	wg.Add(len(sources))

	for _, income := range sources {
		go func(income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mtx.Lock()
				bankBalance += income.Amount
				mtx.Unlock()
				fmt.Printf("On week:%v, you earned: %v, from: %v\n", week, income.Amount, income.Source)
			}
		}(income)
	}

	wg.Wait()

	// print the final balance
	fmt.Printf("Your final earnings:%v\n", bankBalance)
}
