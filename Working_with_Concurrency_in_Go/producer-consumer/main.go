package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int
var pizzaNumber int

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

func makePizza() *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Making pizza #%d, it will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		rnd := rand.Intn(12)
		msg := ""
		success := false
		if rnd < 5 {
			pizzasFailed++
			if rnd <= 2 {
				msg = fmt.Sprintf("*** We ran out of ingredients, pizza for order#%d can not be made!\n", pizzaNumber)
			} else if rnd <= 4 {
				msg = fmt.Sprintf("*** Our chef quit, pizza for order#%d can not be made!\n", pizzaNumber)
			}
		} else {
			pizzasMade++
			msg = fmt.Sprintf("Pizza order#%d is ready!\n", pizzaNumber)
			success = true
		}
		total++

		return &PizzaOrder{pizzaNumber, msg, success}
	}
	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(PizzaMaker *Producer) {

	// keep track of which pizza we are making --> pizzaNumber global variable
	// infintely run
	// create pizzas
	for {
		currentPizza := makePizza()
		select {
		case PizzaMaker.data <- *currentPizza:
		case quitChan := <-PizzaMaker.quit:
			close(PizzaMaker.data)
			close(quitChan)
			return
		}
	}
}

func main() {
	// create and seed the random generator
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// print the starting message
	color.Cyan("Pizza Producing System is working")
	color.Cyan("---------------------------------")

	// create the producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run the consumer

	for order := range pizzaJob.data {
		if order.pizzaNumber <= 10 {
			if order.success {
				color.Green("Order for Pizza#%d in out for delivery", order.pizzaNumber)
			} else {
				color.Red(order.message)
				color.Red("Customer is mad")
			}
		} else {
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel...\n %v", err)
			}
		}
	}

	// print the ending message

	color.Cyan("-----------------------")
	color.Cyan("Pizzeria is closing now")
	color.Cyan("Done for the day...")
	color.Cyan("We made %d pizzas, but failed to make %d pizzas, with %d attempts in total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasMade < 1:
		color.Red("It was a bad day")
	case pizzasMade < 5:
		color.Yellow("It was an ok day")
	case pizzasMade < 8:
		color.Yellow("It was a good day")
	default:
		color.Green("It was a great day")
	}
}
