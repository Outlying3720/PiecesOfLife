package main

import "fmt"

func Fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a + b
		return b
	}
}

func main() {
	fibonacci := Fibonacci()
	for i:=0; i<100; i++ {
		fmt.Println(fibonacci())
	}
}