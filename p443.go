package main

import "fmt"

import "strconv"
func compress(chars []byte) int {
    var result_c = 0
    var now byte = chars[0]
    var count = 0
    var result = 0
    n := len(chars)
	if n <= 1 {
		return n
	}
    for i:=0; i < n; i++ {
        if now != chars[i] {
            if count <= 1 {
                result += 1
				chars[result_c] = byte(now)
                result_c += 1
            } else {
                result += count
				chars[result_c] = byte(now)
                result_c += 1
                var num_s = strconv.Itoa(count)
                for j:=0; j<len(num_s); j++ {
                    chars[result_c] = byte(num_s[j])
                    result_c += 1
                }
            }
            count = 1
            now = chars[i]
            
        } else {
            count += 1
        }
    }
	if count >= 1 {
		if count <= 1 {
			result += 1
			chars[result_c] = byte(now)
			result_c += 1
		} else {
			result += count
			chars[result_c] = byte(now)
			result_c += 1
			var num_s = strconv.Itoa(count)
			for j:=0; j<len(num_s); j++ {
				chars[result_c] = byte(num_s[j])
				result_c += 1
			}
		}
	}
    chars = chars[:result_c]
    return result_c
}

func main() {
	chars := []byte("abc")
	fmt.Print(compress(chars))
	fmt.Print(string(chars))
}