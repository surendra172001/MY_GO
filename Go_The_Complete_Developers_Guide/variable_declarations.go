package main

import (
	"fmt"
)

func VAR_dec() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	//   variable declaration without initial value

	// multiple variable declarations
	var (
		name1 = "surendra"
		name2 = "narendra"
	)

	fmt.Println(name1, "\n", name2)
	// another way of declaring multiple variables
	var p1, p2 = "naveen", "ankur"
	fmt.Println(p1)
	fmt.Println(p2)

	var vint, vstring = 10, "surendra"
	fmt.Println(vint, vstring)

}
