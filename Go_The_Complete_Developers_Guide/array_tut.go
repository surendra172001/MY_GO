package main

import (
	"fmt"
)

func Array_tut() {
	// Declaring array
	// with length
	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	// without length NOTE 3 DOTS
	var arr2 = [...]int{1, 2, 3}
	fmt.Println(arr2)

	// printing array length
	fmt.Println(len(arr2))

	// initialize only specific elements
	var arr3 = [5]int{0: 12, 4: 1}
	fmt.Println(arr3)

	// change values of the elements
	arr3[1] = 56
	fmt.Println(arr3)

	// printing individual elements of the array
	fmt.Println(arr3[1])
}
