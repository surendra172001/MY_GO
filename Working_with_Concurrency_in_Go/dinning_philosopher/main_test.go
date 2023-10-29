package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestDinningPhilosopher(t *testing.T) {
	stdOut := os.Stdout
	os.Stdout, _ = os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Println("Testing the dinning philosopher function")
	dinningPhilosopher()
	if len(eatingOrd) != 15 {
		log.Fatal("There must be 15 elements in eating order")
	}
	os.Stdout = stdOut
}

func TestWithDiffDurations(t *testing.T) {
	stdOut := os.Stdout
	os.Stdout, _ = os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0755)
	myTests := []struct {
		name     string
		duration int
	}{
		{"a second", 1},
		{"twice a second", 2},
		{"no time", 0},
	}

	for _, tInst := range myTests {
		fmt.Printf("Running test %v\n", tInst.name)
		eatDuration = tInst.duration
		eatingOrd = []*Philosopher{}
		dinningPhilosopher()
		if len(eatingOrd) != 15 {
			log.Fatal("There must be 15 elements in eating order")
		}
	}
	os.Stdout = stdOut
}
