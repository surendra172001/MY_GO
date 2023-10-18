package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	filename := os.Args[1]

	// fmt.Printf("The type of args is %T\n", os.Args)

	// fmt.Printf("The length of args is %v\n", len(os.Args))

	// fmt.Println("The input file name:", filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Err: ", err)
	}

	// data := make([]byte, 512)
	// n, err := f.Read(data)
	// if err != nil {
	// 	fmt.Println("Err: ", err)
	// }

	// fmt.Println("The number of characters read: ", n)
	// fmt.Println("The input file data: ")
	// fmt.Println(string(data))

	fmt.Println("///////PRINTING THE FILE CONTENT TO TERMINAL///////")
	io.Copy(os.Stdout, f)
}
