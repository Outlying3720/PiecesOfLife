package main

import "fmt"

func b() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}


func main() {

	fmt.Print(string("这是一个测试"[0]))

	k := 5
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fmt.Println("was <= 5.")
		fmt.Println("was <= 5..")
		fallthrough
	case 6: {
			fmt.Println("was <= 6")
			{
				fmt.Println("was <= 6.")
				fmt.Println("was <= 6..")
			}
		}
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	
	a := 0
	A:
	fmt.Println(a)
	a++
	if a < 15 {
		goto A
	}

	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
    	fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
    	fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}

	b()

	s := []int{1,2,3}
	fmt.Println(s)
	fmt.Println(cap(s))
	// fmt.Println(s[:cap(s) + 1])


	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)

	mp1[1] = []int{1,2,3}
	mp2[2] = &[]int{4,5,6}
	fmt.Println(mp1)
	fmt.Println(mp2)
}