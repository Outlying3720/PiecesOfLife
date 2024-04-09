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

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	l := min(candidates, n)
	r := max(candidates, n-candidates)
	lheap := &IntHeap{}
	rheap := &IntHeap{}
	*lheap = append(*lheap, costs[:l]...)
	*rheap = append(*rheap, costs[r:]...)
	heap.Init(lheap)
	heap.Init(rheap)
	paid := int64(0)
	for i := 0; i < k; i++ {
		if lheap.Len() == 0 {
			paid += int64(heap.Pop(rheap).(int))
			continue
		}
		if rheap.Len() == 0 {
			paid += int64(heap.Pop(lheap).(int))
			continue
		}
		if (*lheap)[0] > (*rheap)[0] {
			paid += int64(heap.Pop(rheap).(int))
			if l < r {
				r--
				heap.Push(rheap, costs[r])
			}
			// fmt.Println("r:", rheap)
		} else {
			paid += int64(heap.Pop(lheap).(int))
			if l < r {
				heap.Push(lheap, costs[l])
				l++
			}
			// fmt.Println("l:", lheap)
		}
	}
	return paid
}

func main() {
	cost := []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}
	fmt.Println(totalCost(cost, 11, 2))
}
