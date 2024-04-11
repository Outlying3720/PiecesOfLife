package quicksort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, st, end int) {
	if st >= end || st < 0 || end >= len(arr) {
		return
	}
	base := arr[st]
	l := st
	r := end
	for l < r {
		for l < r && arr[r] > base {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= base {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = base

	quickSort(arr, st, l-1)
	quickSort(arr, l+1, end)
}

// func main() {
// 	// arr := []int{2, 4, 1, 3, 5, 9, 7, 8, 6}
// 	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
// 	QuickSort(arr)
// 	fmt.Println(arr)
// }
