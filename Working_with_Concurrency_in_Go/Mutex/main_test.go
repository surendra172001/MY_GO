package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = stdOut

	res, _ := io.ReadAll(r)
	out := string(res)

	if !strings.Contains(out, "34320") {
		t.Error("Wrong output")
	}

}
