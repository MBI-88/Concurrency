package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numberofpizzas = 10

var (
	pizzasMade   int
	pizzasFailed int
	total        int
)

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// Producer

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

// Producer end

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberofpizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(pizzaMaker *Producer) {
	var i = 0
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}

}

func consumer(pizzaJob *Producer) {
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberofpizzas {
			if i.success {
				fmt.Printf("Order #%d is out for delivery!\n", i.pizzaNumber)

			} else {
				fmt.Println("The customer is really mad!")
			}
		} else {
			fmt.Println("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				fmt.Println("*** Error closing channel!")
			}
		}
	}

}

func doProducerConsumer() {
	rand.Seed(time.Now().UnixNano())
	pizzaJob := new(Producer)
	pizzaJob.data = make(chan PizzaOrder)
	pizzaJob.quit = make(chan chan error)

	go pizzeria(pizzaJob)

	// Consumer
	consumer(pizzaJob)
	
}
