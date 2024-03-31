package main

import "fmt"
import "runtime"

func main() {
	a := make([]int, 1000000000)
	a[0] = 1
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
}
