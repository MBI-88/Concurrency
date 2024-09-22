package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Exmple race condition test")
	go updateMessage("Goodbye, man!")
	wg.Wait()

	if msg != "Goodbye, man!" {
		t.Error("Incorrect value in msg")
	}

}

func TestUpdateMessage2(t *testing.T) {
	msg = "Hello, cosmos!"
	var mutex sync.Mutex
	wg.Add(2)
	go updateMessage2("Example of fixing race condition", &mutex)
	go updateMessage2("Goodbye, man!", &mutex)
	wg.Wait()

	if msg != "Goodbye, man!" {
		t.Error("Incorrect value in msg")
	}
}

func TestBalance(t *testing.T) {
	stdOut := os.Stdout
	r,w,_ := os.Pipe()

	os.Stdout = w 
	doBalance()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut 
	if !strings.Contains(output, "$34320.00") {
		t.Error("Wrong balance returned")
	}
}