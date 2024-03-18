package main

import "fmt"

func longestSubarray(nums []int) int {
    k := 1
    left := 0
    right := 0
    m := 0
    for _, v := range nums {
        right += 1
		if v == 0 {
            if k > 0 {
                k -= 1
            } else if nums[left] == 0 {
                left += 1
            } else {
                for left < len(nums) && nums[left] == 1 {
                    left += 1
                }
                left += 1
            }
        }
        m = max(m, right - left)
		fmt.Println(nums[left:right])
    }
    return m - 1
}

func main()  {
	nums := []int{0,1,1,1,0,1,1,0,1}
	fmt.Print(longestSubarray(nums))
}