package main

import "fmt"

func SelectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		mi := i
		for j := i; j < n; j++ {
			if arr[mi] > arr[j] {
				mi = j
			}
		}
		arr[mi], arr[i] = arr[i], arr[mi]
	}
}

func main() {
	arr := []int{2, 4, 1, 3, 5, 9, 7, 8, 6}
	SelectSort(arr)
	fmt.Println(arr)
}
