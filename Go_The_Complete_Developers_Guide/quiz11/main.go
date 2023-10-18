// Q4
// package main

// import "fmt"

// func main() {

// 	greeting := "Hi there!"

// 	go (func() {
// 		fmt.Println(greeting)
// 	})()
// }

// Q5
// package main

// func main() {
// 	c := make(chan string)
// 	c <- []byte("Hi there!")
// }

// Q6
package main

func main() {
	c := make(chan string)
	c <- "Hi there!"
}
