package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	for i := 0; i < 10; i++ {
		eatingOrder = []string{}
		eatTime = time.Duration(0)
		sleepTime = time.Duration(0)
		thinkTime = time.Duration(0)
		dine()
		if len(eatingOrder) != 5 {
			t.Errorf("The eating order length %d is not as expected 5", len(eatingOrder))
		}
	}
}

func Test_dineWithVaryingDurations(t *testing.T) {
	var durations = []struct {
		name     string
		duration time.Duration
	}{
		{name: "quarterSecond", duration: 250 * time.Millisecond},
		{name: "halfSecond", duration: 500 * time.Millisecond},
	}

	for _, e := range durations {
		fmt.Println("Test name:", e.name)
		eatingOrder = []string{}
		eatTime = time.Duration(e.duration)
		sleepTime = time.Duration(e.duration)
		thinkTime = time.Duration(e.duration)
		dine()
		if len(eatingOrder) != 5 {
			t.Errorf("The eating order length %d is not as expected 5", len(eatingOrder))
		}
	}
}
