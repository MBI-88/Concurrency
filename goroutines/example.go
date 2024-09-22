package main

import (
	"fmt"
	"sync"
)

func printSomthing(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func doExample() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	wg.Add(len(words))
	for i, x := range words {
		go printSomthing(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
	wg.Add(1)
	printSomthing("This is the second thing to be printed!", &wg)
}
