package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	seatingCapacity = 10
	arrivalRate     = 100
	cutDuration     = 1000 * time.Millisecond
	timeOpen        = 10 * time.Second
)

type BarberShop struct {
	ShopCapacity     int
	HairCutDuration  int
	NumberOfBarbers  int
	BarbersDoneChang chan bool
	ClientsChan      chan string
	Open             bool
}

func (b *BarberShop) addBarber(barber string) {
	b.NumberOfBarbers++
	go func() {
		isSleeping := false
		fmt.Printf("%s goes to the waiting room to check for clients.\n", barber)
		for {
			if len(b.ClientsChan) == 0 {
				fmt.Printf("There is nothing to do, so %s takes a nap\n", barber)
				isSleeping = true
			}
			client, shopOpen := <-b.ClientsChan
			if shopOpen {
				if isSleeping {
					fmt.Printf("%s wakes %s up.\n", client, barber)
					isSleeping = false
				}
				b.cutHair(barber, client)

			} else {
				b.sendBarberHome(barber)
				return
			}
		}
	}()

}

func (b *BarberShop) cutHair(barber, client string) {
	fmt.Printf("%s is cutting %s's hair\n", barber, client)
	time.Sleep(time.Duration(b.HairCutDuration))
	fmt.Printf("%s is finished cutting %s's hair\n", barber, client)
}

func (b *BarberShop) sendBarberHome(barber string) {
	fmt.Printf("%s is going home.\n", barber)
	b.BarbersDoneChang <- true
}

func (b *BarberShop) closeShopForDay() {
	fmt.Println("Closing shop for the day")

	close(b.ClientsChan)
	b.Open = false

	for a := 1; a <= b.NumberOfBarbers; a++ {
		<-b.BarbersDoneChang
	}
	close(b.BarbersDoneChang)
	fmt.Println("The barbershop is now closed for the day, and everyone has gone home")
}

func (b *BarberShop) addClient(client string) {
	fmt.Printf("*** %s arrived!", client)
	if b.Open {
		select {
		case b.ClientsChan <- client:
			fmt.Printf("%s takes seat in the wating room\n", client)
		default:
			fmt.Printf("The waiting room is full, so %s leaves\n", client)
		}
	} else {
		fmt.Printf("The shop is already closed, so %s leaves!\n", client)
	}
}




func doBarber() {

	rand.Seed(time.Now().UnixNano())
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:     seatingCapacity,
		HairCutDuration:  int(cutDuration),
		NumberOfBarbers:  0,
		ClientsChan:      clientChan,
		BarbersDoneChang: doneChan,
		Open:             true,
	}

	fmt.Println("The shop is open for the day!")

	shop.addBarber("Frank")
	shop.addBarber("Gerard")
	shop.addBarber("Milton")
	shop.addBarber("Susan")
	shop.addBarber("Kelly")
	shop.addBarber("Pat")
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	i := 1
	go func() {
		for {
			randomMillseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillseconds)):
				shop.addBarber(fmt.Sprintf("Client #%d\n", i))
				i++
			}
		}
	}()

	time.Sleep(5 * time.Second)

}
