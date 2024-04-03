package main

import "fmt"

func quicksort(nums []int) (ans []int) {
	fmt.Println(nums)
	if len(nums) <= 1 {
		return nums
	}
	base := nums[0]
	l := 0
	r := len(nums) - 1
	for l < r {
		for ; l < r; r-- {
			if base > nums[r] {
				nums[l] = nums[r]
				r--
				break
			}
		}
		for ; l < r; l++ {
			if base < nums[l] {
				nums[r] = nums[l]
				l++
				break
			}
		}
	}
	fmt.Println(r, l)
	nums[l] = base
	if l-1 >= 0 {
		copy(quicksort(nums[:l-1]), ans)
	}
	if l+1 < len(nums) {
		copy(ans, quicksort(nums[l+1:]))
	}
	
	return
}

func main() {
	nums := []int{4,6,3,2,8,7,9,0}
	fmt.Println(quicksort(nums))
}