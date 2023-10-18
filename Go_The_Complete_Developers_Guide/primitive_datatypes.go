package main

import (
	"fmt"
)

func Datatype_tut() {
	// booleans
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	// Integers
	// Signed
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)
	// Unsigned
	var z uint = 500
	var w uint = 18446744073709551615
	fmt.Printf("Type: %T, value: %v\n", z, z)
	fmt.Printf("Type: %T, value: %v\n", w, w)

	// Float
	// float32
	var a float32 = 123.78
	var b float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", a, a)
	fmt.Printf("Type: %T, value: %v\n", b, b)
	// float64
	var c float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v\n", c, c)
	var e = 1.3
	fmt.Printf("Type: %T, value: %v\n", e, e)

	// String
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)

	var ffe float64
	fmt.Printf("Type: %T, value: %f\n", ffe, ffe)

}
