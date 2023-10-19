package main

import "testing"

func TestUpdateMessage(t *testing.T) {
	wg.Add(2)
	go updateMessage("Hello, universe")
	go updateMessage("Hello, cosmos")
	wg.Wait()
	if msg != "Hello, universe" {
		t.Error("Wrong message is being setup")
	}
}
