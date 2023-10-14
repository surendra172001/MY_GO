package main

import (
	"fmt"
)

func Loop_tut() {
	// simple loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
