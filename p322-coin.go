package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	// fmt.Println(coins)
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] == -1 {
				continue
			}
			// dp[i] = min(dp[i], dp[i-coin]+1)
			if dp[i] != -1 {
				// dp[i] = min(dp[i], dp[i-coin]+1)
				if dp[i] < dp[i-coin]+1 {
					// fmt.Println(coin, i, dp[i], i-coin, dp[i-coin]+1)
					dp[i] = dp[i]
				} else {
					dp[i] = dp[i-coin] + 1
				}
			} else {
				dp[i] = dp[i-coin] + 1
			}
			if i == 3126 {
				fmt.Println(coin, i, dp[i], i-coin, dp[i-coin]+1)
			}
			// fmt.Println(dp)
		}
		// for i:=0; i<=amount; i++ { if dp[i]!=-1 { fmt.Print(i, dp[i], "  ") } }
		// fmt.Println("")
		// fmt.Println(dp)
	}

	// if dp[amount] == 9999999 { return -1 }
	return dp[amount]
}

func main() {
	coins := []int{186, 419, 83, 408}
	amount := 3126 //6249

	fmt.Println(coinChange(coins, amount))
}
