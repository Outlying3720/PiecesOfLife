package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

func main() {
	h := &IntHeap{2, 5, 6, 1, 8, 9}
	heap.Init(h)
	heap.Push(h, 4)
	fmt.Println((*h)[0])
	for h.Len() > 0 {
		fmt.Print(heap.Pop(h), " ")
	}
}
