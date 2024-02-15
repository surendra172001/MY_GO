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

/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const totalPizzas = 10

var pizzasMade, pizzasFailed int

type PizzaOrder struct {
	pizzaNumber int
	success     bool
	msg         string
}

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= totalPizzas {
		// compute delay
		delay := rand.Intn(5) + 1
		// print message with pizzanumber and delay
		fmt.Printf("Order#%d is received\n", pizzaNumber)
		// compute a random number to determine the pizza state
		pizzaState := rand.Intn(12) + 1
		success := false
		msg := ""
		if pizzaState < 5 {
			pizzasFailed++
		} else {
			success = true
			pizzasMade++
		}

		fmt.Printf("Preparing Order#%d, it will take %d seconds\n", pizzaNumber, delay)

		time.Sleep(time.Second * time.Duration(delay))

		if pizzaState <= 2 {
			msg = fmt.Sprintf("Order#%d can not completed since ingredients not available", pizzaNumber)
		} else if pizzaState <= 4 {
			msg = fmt.Sprintf("Order#%d can not completed since chef is not available", pizzaNumber)
		} else {
			msg = fmt.Sprintf("Order#%d is complete", pizzaNumber)
		}
		// create pizza and return it
		return &PizzaOrder{pizzaNumber, success, msg}
	}
	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(pizzaMaker *Producer) {
	var i = 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// Printing start message
	fmt.Println("This is a Pizza restaurant")

	// create a producer
	pizzaMaker := Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// start the producer
	go pizzeria(&pizzaMaker)

	// start the consumer
	for i := range pizzaMaker.data {
		if i.pizzaNumber <= totalPizzas {
			fmt.Println(i.msg)
			if i.success {
				fmt.Printf("Order#%d is out for delivery\n", i.pizzaNumber)
			} else {
				fmt.Println("Customer is mad...")
			}
		} else {
			err := pizzaMaker.Close()
			if err != nil {
				fmt.Printf("There is some error while closing the program - %v\n", err)
			}
			close(pizzaMaker.quit)
		}
	}

	// print final stats and messages
	fmt.Printf("Pizzas made - %d\n", pizzasMade)
	fmt.Printf("Pizzas failed - %d\n", pizzasFailed)
	fmt.Printf("Total Orders - %d\n", pizzasMade+pizzasFailed)
	if pizzasFailed < 2 {
		fmt.Println("It was a great day at work, The restaurant has pretty high standards")
	} else if pizzasFailed < 4 {
		fmt.Println("It was a good day at work, The restaurant has decent standards")
	} else {
		fmt.Println("It was a disastorous day at work, The restaurnat has poor standards")
	}
}

*/
