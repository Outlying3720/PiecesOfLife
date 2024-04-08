package main

import "fmt"

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{2, 4, 1, 3, 5, 9, 7, 8, 6}
	BubbleSort(arr)
	fmt.Println(arr)
}
