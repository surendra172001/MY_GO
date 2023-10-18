package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func readFwriteF() {
	// stdIn := os.Stdin
	stdOut := os.Stdout
	// os.Stdin, _ = os.Open("./input.txt")
	inputFile, err := os.Open("./input.txt")
	os.Stdout, _ = os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	res := make([]byte, 200)

	inputFile.Read(res)

	res = bytes.Trim(res, "\x00")

	fmt.Println("output from file", string(bytes.Trim(res, "\x00")))
	// fmt.Printf("Len: %v, output: %v\n", len(res), string(bytes.Trim(res, "\x00")))

	// os.Stdin = stdIn
	os.Stdout = stdOut
}

func stdIO() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter text (Ctrl+D to finish):")

	// Read lines one by one
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("You entered:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
	}
}

func main() {
	readFwriteF()

	stdIO()
}
