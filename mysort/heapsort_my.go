package main

import "fmt"

type Heap []int

func (h Heap) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h *Heap) Swap(a, b int) {
	(*h)[a], (*h)[b] = (*h)[b], (*h)[a]
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) up(idx int) {
	p := (idx - 1) / 2
	if p < 0 {
		return
	}

	if h.Less(idx, p) {
		h.Swap(p, idx)
		h.up(p)
	}
}

func (h Heap) down(idx int) {
	l := idx*2 + 1
	r := idx*2 + 2

	if l >= h.Len() {
		return
	}
	minest_child_idx := l
	if r < h.Len() && h.Less(r, l) {
		minest_child_idx = r
	}
	if h.Less(minest_child_idx, idx) {
		h.Swap(idx, minest_child_idx)
		h.down(minest_child_idx)
	}
}

func (h Heap) Init() {
	for i := (h.Len() - 1) / 2; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Heap) Push(num int) {
	*h = append(*h, num)
	h.up(h.Len() - 1)
}

func (h *Heap) Pop() (int, bool) {
	return h.Remove(0)
}

func (h *Heap) Remove(idx int) (int, bool) {
	if h.Len() == 0 || idx >= h.Len() {
		return 0, false
	}
	num := (*h)[idx]
	(*h)[idx] = (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	h.down(idx)
	return num, true
}

func HeapSort(arr []int) {
	heap := Heap(arr)
	heap.Init()
	ans := make([]int, len(arr))
	for i := range heap {
		ans[i], _ = heap.Pop()
	}
	copy(arr, ans)
}

func main() {
	arr := []int{2, 4, 1, 3, 5, 9, 7, 8, 6}
	HeapSort(arr)
	fmt.Println(arr)

	// arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	// heap := Heap(arr)
	// heap.Init()
	// fmt.Println(heap)
	// // heap.Push(5)
	// // heap.Push(50)
	// // heap.Push(15)
	// // heap.Push(25)
	// // heap.Push(2)
	// // heap.Push(4)
	// // heap.Push(7)
	// // heap.Push(35)
	// // heap.Push(30)
	// // heap.Push(17)
	// // heap.Push(12)

	// fmt.Println(heap.Remove(5))
	// fmt.Println(heap.Remove(2))
	// fmt.Println(heap.Remove(2))
	// fmt.Println(heap.Remove(8))

	// for i := heap.Len(); i > 0; i-- {
	// 	fmt.Println(heap.Pop())
	// }
}
