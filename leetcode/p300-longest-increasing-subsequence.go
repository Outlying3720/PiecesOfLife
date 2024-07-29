package main

import "fmt"

func lengthOfLIS(nums []int) int {
	len2min := map[int]int{}
	n := len(nums)
	dp := make([]int, n)

	dp[0] = 1
	len2min[1] = nums[0]

	for i := 1; i < n; i++ {
		m := 1
		for j := dp[i-1]; j > 0; j-- {
			if nums[i] > len2min[j] {
				m = max(m, j+1)
			}
		}
		if m != 1 || nums[i] < len2min[m] {
			len2min[m] = nums[i]
		}

		dp[i] = max(dp[i-1], m)
	}

	fmt.Println(dp)
	fmt.Println(len2min)

	return dp[n-1]
}

func main() {
	nums := []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(nums))
}
