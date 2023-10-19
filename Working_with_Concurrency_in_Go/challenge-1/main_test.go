package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// func TestUpdateMessage(t *testing.T) {

// 	wg.Add(1)
// 	go updateMessage("Dil maange more")
// 	wg.Wait()

// 	if !strings.Contains(msg, "Dil") {
// 		t.Error("The message must contain Dil")
// 	}
// }

// func TestPrintMessage(t *testing.T) {
// 	stdOut := os.Stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	msg = "Dil Maange More"
// 	printMessage()

// 	res, _ := io.ReadAll(r)
// 	output := string(res)

// 	if !strings.Contains(output, "Dil") {
// 		t.Error("The message must contain Dil")
// 	}

// 	os.Stdout = stdOut
// }

func TestUpdateMessage(t *testing.T) {
	wg.Add(1)
	updateMessage("Hi there surendra")
	wg.Wait()
	if !strings.Contains(msg, "there") {
		t.Error("Message should contain there")
	}
}

func TestPrintMessage(t *testing.T) {
	msg = "surendra"
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printMessage()
	_ = w.Close()
	os.Stdout = stdOut
	res, _ := io.ReadAll(r)
	output := string(res)
	if !strings.Contains(output, "surendra") {
		t.Error("Print message must contain surendra")
	}
}

func TestMain(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	os.Stdout = stdOut
	res, _ := io.ReadAll(r)
	output := string(res)

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Print message must contain Hello! universe")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Print message must contain Hello! cosmos")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Error("Print message must contain Hello! world")
	}
}
