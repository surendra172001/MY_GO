package main

import (
	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func pizzeria(PizzaMaker *Producer) {
	// infintely run
	// create pizzas
	for {

	}
}

func main() {
	// create and seed the random generator
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// print the starting message
	color.Red("Pizza Producing System is working")
	color.Red("---------------------------------")

	// create the producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run the consumer in the background

	// print the ending message
}
