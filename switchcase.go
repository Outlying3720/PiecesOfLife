package main

import "fmt"
import "strings"

func reverseVowels(s string) string {
    n := len(s)
    t := []byte(s)
    var tmp [300000]byte
    var tmp_n = 0
    for i:=n-1; i>=0; i-- {
        // fmt.Print("%c", s[i])
        if strings.Contains("aeiouAEIOU", string(s[i])) {
            tmp[tmp_n] = s[i]
            tmp_n += 1
            // fmt.Print(s[i])
        }
    }
    tmp_n = 0
    for i:=0; i<n; i++ {
        if strings.Contains("aeiouAEIOU", string(s[i])) {
            t[i] = tmp[tmp_n]
            tmp_n += 1
        }
    }
    return string(t)
}

func main() {
    fmt.Print(reverseVowels("aeiou"))
}
