package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func searchFile(dir string, lookFor string, ch chan string) {
	// log.Println("[SEARCHING] ", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// log.Println(filepath.Join(dir, file.Name()), file.IsDir())
		if file.Name() == lookFor {
			ch <- "[FOUND] " + filepath.Join(dir, file.Name())
			return
		}
	}
	ch <- "[NOT FOUND] " + dir
}

func searchFileDemo() {
	fileName := "main.go"
	ch := make(chan string)
	go searchFile("bills", fileName, ch)
	go searchFile("cards", fileName, ch)
	go searchFile("coupa", fileName, ch)
	go searchFile("externalauth", fileName, ch)

	for val := range ch {
		fmt.Printf("%T, %v\n", val, val)
	}

	// flag := true
	// for i := 0; i < 4 && flag; i++ {
	// 	select {
	// 	case val, ok := <-ch:
	// 		if ok {
	// 			fmt.Println(val)
	// 		} else {
	// 			flag = false
	// 		}
	// 	}
	// }
}
