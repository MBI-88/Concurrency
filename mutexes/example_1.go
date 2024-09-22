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

func doExample() {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	wg.Wait()

	fmt.Println(msg)
}