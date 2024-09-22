package main

import (
	"fmt"
	"sync"
)


type Income struct {
	Source string
	Amount int
}

func doBalance() {
	var (
		bankBalance int 
		mutex sync.Mutex
	)

	fmt.Printf("Initial account balance: $%d.00\n", bankBalance)
	incomes := []Income{
		{Source: "Main job",Amount: 500},
		{Source: "Gifts",Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func (i int,income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mutex.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp 
				mutex.Unlock()
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week,income.Amount,income.Source )
			}
		}(i,income)
	}
	wg.Wait()
	fmt.Printf("Final banck balance: $%d.00",bankBalance)


}