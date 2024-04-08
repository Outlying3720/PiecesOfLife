package main

import (
	"container/heap"
	"fmt"
	"slices"
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

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	ids := make([]int, len(nums1))
	for i := range ids {
		ids[i] = i
	}
	slices.SortFunc(ids, func(i, j int) int { return nums2[j] - nums2[i] })

	h := make(IntHeap, k)
	sum := 0
	for i, idx := range ids[:k] {
		sum += nums1[idx]
		h[i] = nums1[idx]
	}
	heap.Init(&h)

	ans := sum * nums2[ids[k-1]]
	for _, i := range ids[k:] {
		x := nums1[i]
		if x > h[0] {
			sum += x - h[0]
			h[0] = x
			heap.Fix(&h, 0)
			ans = max(ans, sum*nums2[i])
		}
	}
	return int64(ans)
}
func main() {
	nums1 := []int{1, 3, 3, 2}
	nums2 := []int{2, 1, 3, 4}
	fmt.Println(maxScore(nums1, nums2, 3))
}
