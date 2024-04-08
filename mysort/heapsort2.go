package main

import "fmt"

type Heap struct {
	arr  []int
	size int
}

func (h Heap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h Heap) Less(i, j int) bool {
	return h.arr[i] < h.arr[j]
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
	if len(h.arr) < h.size {
		h.arr = append(h.arr, x)
		h.up(len(h.arr) - 1)
	} else if x > h.arr[0] {
		h.arr[0] = x
		h.down(0)
	}

}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1
		r := 2*i + 2
		if l >= len(h.arr) {
			break
		}
		// j 为子节点中最小的
		j := l
		if r < len(h.arr) && h.Less(r, l) {
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

func findKthLargest(arr []int, size int) int {
	h := Heap{size: size}
	for _, num := range arr {
		h.Push(num)
	}

	return h.arr[0]
}

func main() {
	arr := []int{5, 6, 3, 7, 8, 9, 1, 2, 4, 9}
	fmt.Println(findKthLargest(arr, 6))
}
