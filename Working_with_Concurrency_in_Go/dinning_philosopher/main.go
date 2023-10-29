package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	hunger        = 3
	forkCnt       = 5
	eatDuration   = 3
	thinkDuration = 1
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var wg sync.WaitGroup
var seated sync.WaitGroup

var eatOrdLock sync.Mutex
var eatingOrd []*Philosopher

func printEatOrd() {
	fmt.Print("Printing eating order\n")

	for i, p := range eatingOrd {
		fmt.Printf("%dth philosopher is %v\n", i, p.name)
	}
}

func dine(philosopher *Philosopher, forks map[int]*sync.Mutex) {
	defer wg.Done()
	seated.Done()
	fmt.Printf("Seated: %v\n", philosopher.name)
	seated.Wait()
	for i := hunger; i > 0; i-- {
		// pick the larger number fork first
		if philosopher.leftFork > philosopher.rightFork {
			fmt.Printf("Left Fork pickup: %v\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("Right Fork pickup: %v\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
		} else {
			fmt.Printf("Right Fork pickup: %v\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("Left Fork pickup: %v\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
		}
		fmt.Printf("Eating: %v\n", philosopher.name)
		eatOrdLock.Lock()
		eatingOrd = append(eatingOrd, philosopher)
		eatOrdLock.Unlock()
		time.Sleep(time.Duration(eatDuration))
		fmt.Printf("Thinking: %v\n", philosopher.name)
		time.Sleep(time.Duration(thinkDuration))

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
	}
	fmt.Printf("Philosopher %v is satisfed\n", philosopher.name)
	fmt.Printf("Philosopher %v has left the table\n", philosopher.name)
}

func dinningPhilosopher() {

	// create philosopher
	philosophers := []Philosopher{
		{"Aristotle", 0, 1},
		{"Socrates", 1, 2},
		{"Hercules", 2, 3},
		{"Neuman", 3, 4},
		{"Aryabhatt", 4, 0},
	}

	// populating map of the mutexes
	forks := make(map[int]*sync.Mutex)

	for i := 0; i < forkCnt; i++ {
		forks[i] = &sync.Mutex{}
	}

	wg.Add(len(philosophers))
	seated.Add(len(philosophers))
	for i := 0; i < len(philosophers); i++ {
		go dine(&philosophers[i], forks)
	}

	wg.Wait()

}

func main() {

	fmt.Print("The program output is going in output.txt file\n")

	stdOut := os.Stdout

	os.Stdout, _ = os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0755)

	// print starting message
	fmt.Println("Dinning started philosophers will eat now")
	fmt.Println("-----------------------------------------")

	dinningPhilosopher()

	fmt.Print("----------------------------------------\n")
	fmt.Print("Philosophers have completed their dinner\n")
	printEatOrd()
	os.Stdout = stdOut
}
