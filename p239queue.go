package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) (ans []int) {
	queue := []int{}
	n := len(nums)
	push := func(i int) {
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	for i := 0; i < min(k, n); i++ {
		push(i)
	}
	ans = append(ans, nums[queue[0]])
	for i := k; i < n; i++ {
		push(i)

		for queue[0] <= i-k {
			queue = queue[1:]
		}

		ans = append(ans, nums[queue[0]])
	}
	return
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
