package main

import (
	"fmt"
	"sync"
	"time"
)

// creating philosophers
type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

/*
		f0		f4
			P0
		p1		p4
	f1				f3
		p2		p3
			f2

P0 - Plato
P1 - Socrates
P2 - Aristotle
p3 - Pascal
p4 - Locke

forks - f0, f1, f2, f3, f4
*/

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// declaration of time durations
var hungry = 3
var eatTime = time.Second
var sleepTime = time.Duration(3) * time.Second
var thinkTime = time.Second

var orderMutex sync.Mutex
var eatingOrder []string

// create waitgroup to wait for dining process
// create waitgroup to wait for seating at the table
// add count to the waitgroup for dinning and seating
// create map of mutexes for forks
// call the function for dinning process
// wait for the completion of dining process
func dine() {
	var wg sync.WaitGroup
	wg.Add(len(philosophers))

	var seated sync.WaitGroup
	seated.Add(len(philosophers))

	var forks = map[int]*sync.Mutex{}

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go dinePhilosopher(philosophers[i], forks, &wg, &seated)
	}

	wg.Wait()

}

// Decrease the counter for dining process
// print the statement for seating
// Decrease the counter for seating
// wait for all the philosophers to be seated
// run a for loop to eat hungry number of times
// in each iteration pick 2 forks in increasing order of their id
// start eating
// drop the picked forks
// start thinking
// print the message - philosopher is satisfied and leaving...
func dinePhilosopher(philosopher Philosopher, forks map[int]*sync.Mutex, wg *sync.WaitGroup, seated *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s is seated on the table\n", philosopher.name)
	seated.Done()
	seated.Wait()

	for i := 0; i < hungry; i++ {
		if philosopher.leftFork < philosopher.rightFork {
			forks[philosopher.leftFork].Lock()
			printMsg(i, philosopher.name, "picked up left fork")
			forks[philosopher.rightFork].Lock()
			printMsg(i, philosopher.name, "picked up right fork")
		} else {
			forks[philosopher.rightFork].Lock()
			printMsg(i, philosopher.name, "picked up right fork")
			forks[philosopher.leftFork].Lock()
			printMsg(i, philosopher.name, "picked up left fork")
		}
		printMsg(i, philosopher.name, "is eating")
		time.Sleep(eatTime)
		printMsg(i, philosopher.name, "is thinking")
		time.Sleep(thinkTime)
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
		printMsg(i, philosopher.name, "is sleeping")
		time.Sleep(sleepTime)
	}

	printMsg(-1, philosopher.name, "is satisfied and leaving")
	orderMutex.Lock()
	eatingOrder = append(eatingOrder, philosopher.name)
	orderMutex.Unlock()
}

func printMsg(iteration int, name string, msg string) {
	fmt.Printf("%d - %s %s...\n", iteration, name, msg)
}

func main() {
	eatTime = time.Duration(0)
	sleepTime = time.Duration(0)
	thinkTime = time.Duration(0)
	// print the start message
	fmt.Println("The dinning philosophers problem")

	// start the process
	fmt.Println("The table is empty")

	dine()

	// end the process
	fmt.Println("The table is empty")
}
