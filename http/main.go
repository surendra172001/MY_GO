package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com")

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s", body)

	lg := logWriter{}
	n, err := lg.Write(body)

	fmt.Println("Number of bytes written to the stdout", n)

	if err != nil {
		fmt.Println("err", err)
	}

}

// Writing our own Write function for implementing Writer interface
type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
