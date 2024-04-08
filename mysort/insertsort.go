package main

import "fmt"

func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		slct := arr[i]
		j := i - 1
		for j >= 0 && slct < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = slct
	}
}

func main() {
	arr := []int{2, 4, 1, 3, 5, 9, 7, 8, 6}
	InsertSort(arr)
	fmt.Println(arr)
}
