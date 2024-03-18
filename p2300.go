package main

import "fmt"
import "slices"

func successfulPairs(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)
    result := make([]int, len(spells))
    for idx, spell := range spells {
        left, right := 0, len(potions)
        for left < right {
            mid := (left + right) / 2
            if int64(potions[mid]) * int64(spell) >= success {
                right = mid
            } else {
                left = mid + 1
            }
        }
        if left == right {
            result[idx] = len(potions) - left
        }
    }
    return result
}

func main() {
	spells := []int{3,1,2}
	potions := []int{8,5,8}
	success := int64(16)
	fmt.Println(successfulPairs(spells, potions, success))
}