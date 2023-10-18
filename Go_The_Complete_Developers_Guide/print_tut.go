package main

import (
	"fmt"
)

func Print_tut() {
	//   print function adds a space between 2 args if none of them is a string
	var v1, v2 = 10, 20
	fmt.Print(v1, v2, "\n")

	// using println
	fmt.Println(v1, v2)

	// using printf
	var v3 = "mahendra"
	fmt.Printf("V3 has value: %v and type: %T\n", v3, v3)
	fmt.Printf("V2 has value: %v and type: %T\n", v2, v2)

	// fmt.Print(v1, v3)

	// formatting integers
	fmt.Printf("%-4d %4d\n", v1, v2)

	var i = 15

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)

	//   string formatting
	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	// float formatting verbs
	var f = 3.141

	fmt.Printf("%e\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%6.2f\n", f)
	fmt.Printf("%g\n", f)
}
