package main

import (
	"fmt"
	"math/rand"
)

func quickSelect3Ways(nums []int, k int) int {
	swap := func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }
	var sort func(l, r int) int
	sort = func(l, r int) int {
		if l == r {
			fmt.Println(l, r)
			return nums[l]
		}
		swap(l, rand.Intn(r-l+1)+l)
		v := nums[l]
		lt := l
		gt := r + 1
		i := l + 1
		for i < gt {
			if nums[i] < v {
				swap(i, lt+1)
				lt++
				i++
			} else if nums[i] > v {
				swap(i, gt-1)
				gt--
			} else {
				i++
			}
		}
		swap(l, lt)
		fmt.Println(nums, lt, gt, len(nums)-k)
		if lt <= len(nums)-k && gt > len(nums)-k {
			fmt.Println(lt, gt, nums[lt], nums[gt])
			return nums[lt]
		}
		if lt > len(nums)-k {
			return sort(l, lt-1)
		} else {
			return sort(gt, r)
		}
	}

	return sort(0, len(nums)-1)
}

func findKthLargest(nums []int, k int) int {
	return quickSelect3Ways(nums, k)
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
