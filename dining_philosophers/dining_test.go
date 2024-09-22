package main

import (
	"testing"
	"time"
)

func TestDining(t *testing.T) {
	for i := 0; i < 10; i++ {
		orderfinished = []string{}
		dine()
		if len(orderfinished) != 5 {
			t.Errorf("Incorrect length of slice; expected 5 but got %d",len(orderfinished))
		}
	}
}

func TestDelay(t *testing.T) {
	var theTests = []struct{
		name string
		delay time.Duration
	}{
		{"zero delay",time.Second * 0},
		{"quarter second delay",time.Millisecond * 250},
		{"half second delay",time.Millisecond * 500},
	}

	for _,e := range theTests {
		orderfinished = []string{}

		eatTime = e.delay
		thinkTime = e.delay
		dine()
		if len(orderfinished) != 5 {
			t.Errorf("Incorrect length of slice; expected 5 but got %d",len(orderfinished))
		}
	}
}