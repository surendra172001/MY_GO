package main

import "fmt"

func updateName(x *string) {
	*x = "narendra"
}

func Pointer_tut() {
	var name = "surendra"
	fmt.Println(name)
	updateName(&name)
	fmt.Println(name)
}
