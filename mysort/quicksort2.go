package main

import "fmt"

func quicksort(nums []int, begin int, end int) {
	fmt.Println(nums)
	if len(nums) == 0 {
		return
	}
    if begin >= end {
        return
    }
	base := nums[begin]
	l := begin
	r := end
	for {
		for ; l < r; r-- {
			if base > nums[r] {
				nums[l] = nums[r]
				break
			}
		}
		for ; l < r; l++ {
			if base < nums[l] {
				nums[r] = nums[l]
				break
			}
		}
		if l >= r {
			nums[l] = base
			quicksort(nums, begin, l-1)
			quicksort(nums, l+1, end)
			return
		}
	}
	
}

func main() {
	nums := []int{9,8,7,6,5,4,3,2,1}
	quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}