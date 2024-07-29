package main

import (
	"container/heap"
	"fmt"
)

type IntWithIdx struct{ val, idx int }

type IntWithIdxHeap []IntWithIdx

func (h IntWithIdxHeap) Len() int           { return len(h) }
func (h IntWithIdxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h IntWithIdxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntWithIdxHeap) Push(x interface{}) {
	*h = append(*h, x.(IntWithIdx))
}

func (h *IntWithIdxHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	maxheap := &IntWithIdxHeap{}
	n := len(nums)
	for i := 0; i < min(k, n); i++ {
		item := IntWithIdx{val: nums[i], idx: i}
		heap.Push(maxheap, item)
	}
	ans = append(ans, (*maxheap)[0].val)
	for i := k; i < n; i++ {
		item := IntWithIdx{val: nums[i], idx: i}
		heap.Push(maxheap, item)
		for (*maxheap)[0].idx <= i-k {
			heap.Pop(maxheap)
		}
		ans = append(ans, (*maxheap)[0].val)
	}
	return
}

func main() {
	nums := []int{1}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
