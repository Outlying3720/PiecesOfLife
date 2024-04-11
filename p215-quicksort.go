package main

import "fmt"

// so slow
func quickSort(arr []int, st, end, k int) int {
	if st >= end || st < 0 || end >= len(arr) {
		return arr[st]
	}
	base := arr[st]
	l := st - 1
	r := end + 1
	for l < r {
		for l++; arr[l] < base; l++ {
		}
		for r--; arr[r] > base; r-- {
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[l] = base
	if l != k {
		if l > k {
			return quickSort(arr, st, l-1, k)
		} else {
			return quickSort(arr, l+1, end, k)
		}
	}
	return base
}

// // something wrong
// func quickSort(arr []int, st, end, k int) int {
// 	if st == end {
// 		return arr[st]
// 	}
// 	base := arr[st]
// 	l := st - 1
// 	r := end + 1
// 	for l < r {
// 		for l++; arr[l] < base; l++ {
// 		}
// 		for r--; arr[r] > base; r-- {
// 		}
// 		if l < r {
// 			arr[l], arr[r] = arr[r], arr[l]
// 		}
// 	}
// 	arr[l] = base
// 	if r >= k {
// 		return quickSort(arr, st, r, k)
// 	} else {
// 		return quickSort(arr, r+1, end, k)
// 	}
// }

func findKthLargest(arr []int, k int) int {
	return quickSort(arr, 0, len(arr)-1, len(arr)-k)
}

func main() {
	arr := []int{-1, -1}
	// 1,2,2,3,3,4,5,5,6
	fmt.Println(findKthLargest(arr, 4))
}
