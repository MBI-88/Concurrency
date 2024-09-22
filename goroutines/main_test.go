package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"

)

func TestPrintSomething(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)
	go printSomthing("epsilon", &wg)

	wg.Wait()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	ouput := string(result)

	os.Stdout = stdOut
	if !strings.Contains(ouput, "epsilon") {
		t.Errorf("Expected to find epsilon, but it is not there")
	}

}

func TestUpdateMessage(t *testing.T) {
	var messages = "Hello, world test!"

	wg.Add(1)
	go updateMessage(messages)
	wg.Wait()

	if msg != messages {
		t.Errorf("%s != %s", msg,messages)
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	wg.Add(1)
	go updateMessage("Hello test print message")
	wg.Wait()
	printMessage()

	_ = w.Close()
	output, _ := io.ReadAll(r)
	resl := string(output)
	os.Stdout = stdOut
	if !strings.Contains(resl,"Hello test print message" ) {
		t.Errorf("The result is not iqual to input")
	}
}