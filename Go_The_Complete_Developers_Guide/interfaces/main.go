package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreetings() string {
	return "Hi There!"
}

func (spanishBot) getGreetings() string {
	return "Hola!"
}

type bot interface {
	getGreetings() string
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}
