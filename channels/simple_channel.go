package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func doChannel() {
	var userInput string
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press Enter (enter Q to quit)")

	for {
		fmt.Print("-> ")
		_, _ = fmt.Scanln(&userInput)
		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		resp := <-pong
		fmt.Printf("Response: %s\n", resp)

	}
	fmt.Println("All done. Closing channels")
	close(ping)
	close(pong)
}
