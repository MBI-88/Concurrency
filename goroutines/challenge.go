package main

import (
	"fmt"
	"sync"
)

var (
	msg string
	wg sync.WaitGroup
)

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func doChallenge() {
	/*
	 challenge: modify this code so that the calls to updateMessage() run as goroutines, and
	 implement wait groups so that the program runs properly, and prints out three different
	 messages.
	 Then, write test for all three functions in this program: updateMessage(),printMessage(), and main()
	*/
	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()
}
