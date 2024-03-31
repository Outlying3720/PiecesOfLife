package main

import (
	"fmt"
	"time"
)

func upper(out chan float64) {
	for {
		out <- 4.0
		out <- -4.0
	}
}

func lower(out chan float64) {
	for i:=0;;i++{
		out <- 2.0 * float64(i) + 1.0
	}
}

func calcPi(out chan float64) {
	upperchan, lowerchan := make(chan float64), make(chan float64)
	go upper(upperchan)
	go lower(lowerchan)
	result := 0.0
	for {
		result += <-upperchan / <-lowerchan
		out <- result
	}
}

func main() {
	start := time.Now()
	out := make(chan float64)
	go calcPi(out)
	for i := 0; i < 1e6; i++ {
		<-out
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	fmt.Println(<-out)
}