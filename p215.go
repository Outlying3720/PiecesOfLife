package main

import "fmt"

// UNDONE

import (
    "container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
	h := make(IntHeap, k)
	copy(h, nums[:k])
	heap.Init(&h)
	for i := k; i < len(nums); i++ {
		if x:=nums[i]; x < h[0] {
			fmt.Println(h)
			heap.Pop(&h)
			heap.Push(&h, x)
		}
	}
	fmt.Println(h)
	return h[0]
}

func main() {
	nums := []int{3,2,1,5,6,4}
	k:=2
	fmt.Println(findKthLargest(nums, k))
}