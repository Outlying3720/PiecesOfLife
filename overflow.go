package main

import "fmt"

func main() {
	a := byte(1)
	for i := 0; i < 300; i++ {
		a = a + 1
		fmt.Printf("%c", a)
	}
}
