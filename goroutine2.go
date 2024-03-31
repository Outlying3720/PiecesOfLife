package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(2e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	time.Sleep(3e8)
	ch <- "Tripoli"
	time.Sleep(3e8)
	ch <- "London"
	time.Sleep(3e8)
	ch <- "Beijing"
	time.Sleep(3e8)
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}