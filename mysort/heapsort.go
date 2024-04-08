package main

import "fmt"

type Heap []int

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

// 将当前位置与父节点比较、交换
func (h Heap) up(i int) {
	for {
		parenti := (i - 1) / 2
		if i == parenti || h.Less(parenti, i) {
			break
		}
		h.Swap(parenti, i)
		i = parenti
	}
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1
		r := 2*i + 2
		if l >= len(h) {
			break
		}
		// j 为子节点中最小的
		j := l
		if r < len(h) && h.Less(r, l) {
			j = r
		}
		// 如果父节点(当前)比子节点小退出
		if h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func (h *Heap) Remove(i int) (int, bool) {
	n := len(*h) - 1
	if i < 0 || i > n {
		return 0, false
	}
	h.Swap(i, n) // 删除时将要删除的值移到最后
	x := (*h)[n]
	*h = (*h)[0:n]
	if i == (i-1)/2 || (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else {
		h.up(i)
	}
	return x, true
}

func (h *Heap) Pop() int {
	x, _ := h.Remove(0)
	return x
}

func BuildHeap(arr []int) Heap {
	h := Heap(arr)
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}

	return h
}

func HeapSort(arr []int) {
	// 创建堆
	heap := BuildHeap(arr)
	var sortedArr []int
	for len(heap) > 0 {
		sortedArr = append(sortedArr, heap.Pop())
	}

	fmt.Println(sortedArr)
}

func main() {
	arr := []int{5, 6, 3, 7, 8, 9, 1, 2, 4, 9}
	heap := BuildHeap(arr)
	fmt.Println(heap)
	HeapSort(arr)
}
