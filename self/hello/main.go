package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("surendra")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	msgs, err := greetings.Hellos([]string{"surendra", "narendra", "shivansh", "rajkumar"})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range msgs {
		fmt.Println(v)
	}
}

// func main() {
// 	fmt.Println(quote.Go())
// 	fmt.Println(greetings.Hello("Surendra"))
// 	fmt.Println("Welcome to Todo List App")
// }
