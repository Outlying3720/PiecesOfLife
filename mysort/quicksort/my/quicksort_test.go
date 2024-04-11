package quicksort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		list = append(list, rand.Intn(10000))
	}
	QuickSort(list)
	fmt.Println(list)
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := make([]int, 0, 1000)
		for i := 0; i < 1000; i++ {
			list = append(list, rand.Intn(10000))
		}
		QuickSort(list)
	}
}

func BenchmarkQuickSortGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := make([]int, 0, 1000)
		for i := 0; i < 1000; i++ {
			list = append(list, rand.Intn(10000))
		}
		QuickSortGo(list)
	}
}
