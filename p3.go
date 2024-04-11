package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	hashset := make(map[byte]int)
	left := 0
	m := 0
	tmp := 0
	for cur := range s {
		c := s[cur]
		if idx, exist := hashset[c]; exist {
			m = max(m, tmp)
			i := left
			for ; i <= idx; i++ {
				delete(hashset, s[i])
				tmp -= 1
				if s[i] == c {
					break
				}
			}
			left = idx + 1
		}
		tmp += 1
		hashset[c] = cur
	}
	m = max(m, tmp)
	return m
}

func main() {
	s := "aabaabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
