package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan int)
	result := 0
	start := time.Now()
	for i := 0; i <= 25; i++ {
		go go_fibonacci(i, out)
	}
	for i := 0; i <= 25; i++ {
		result = <- out
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func go_fibonacci(n int, out chan<- int) {
	out <- fibonacci(n)
}

func fibonacci(n int) (res int){
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
