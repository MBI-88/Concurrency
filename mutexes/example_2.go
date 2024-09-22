package main

import (
	"fmt"
	"sync"
)



func updateMessage2(s string,m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func doExample2() {
	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage2("Hello, world!",&mutex)
	go updateMessage2("Hello, cosmos!",&mutex)
	wg.Wait()
	fmt.Println(msg)

}