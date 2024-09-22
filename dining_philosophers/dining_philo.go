package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var (
	hunger      = 3
	eatTime     = 0 * time.Second // for testing value 0
	thinkTime   = 0 * time.Second
	philosopers = []Philosopher{
		{name: "Plato", leftFork: 4, rightFork: 0},
		{name: "Socrates", leftFork: 0, rightFork: 1},
		{name: "Aristotle", leftFork: 1, rightFork: 2},
		{name: "Pascal", leftFork: 2, rightFork: 3},
		{name: "Locke", leftFork: 3, rightFork: 4}}
	forks = make(map[int]*sync.Mutex)
	order  = sync.Mutex{}
	orderfinished = []string{}
)

func diningProblem(philosoper Philosopher, wg *sync.WaitGroup, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table\n", philosoper.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {
		if philosoper.leftFork > philosoper.rightFork {
			forks[philosoper.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosoper.name)
			forks[philosoper.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosoper.name)
		} else {
			forks[philosoper.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosoper.name)
			forks[philosoper.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosoper.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosoper.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosoper.name)
		time.Sleep(thinkTime)

		forks[philosoper.leftFork].Unlock()
		forks[philosoper.rightFork].Unlock()
		fmt.Printf("\t%s put down the forks.\n", philosoper.name)
	}

	fmt.Println(philosoper.name, "is satisified")
	fmt.Println(philosoper.name, "left the table")
	order.Lock()
	orderfinished = append(orderfinished, philosoper.name)
	order.Unlock()

}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosopers))
	seated := &sync.WaitGroup{}
	seated.Add(len(philosopers))

	for i := 0; i < len(philosopers); i++ {
		forks[i] = &sync.Mutex{}
	}
	for i := 0; i < len(philosopers); i++ {
		go diningProblem(philosopers[i], wg, seated)    
	}

	wg.Wait()
	fmt.Printf("Order finished: %s.\n", strings.Join(orderfinished, ", "))
	fmt.Println("The table is empty")

}
