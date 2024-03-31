package main

import (  
    "fmt"
    "time"
)

func forState1(){
    data := []string{"one","two","three"}

    for _,v := range data {
        go func() {
            fmt.Println(v)
        }()
    }
    time.Sleep(3 * time.Second)    //goroutines print: three, three, three

	// actually now fixed? 2024.03.28

    for _,v := range data {
        vcopy := v // 使用临时变量
        go func() {
            fmt.Println(vcopy)
        }()
    }
    time.Sleep(3 * time.Second)    //goroutines print: one, two, three

    for _,v := range data {
        go func(in string) {
            fmt.Println(in)
        }(v)
    }
    time.Sleep(3 * time.Second)    //goroutines print: one, two, three
}

func main() {  
    forState1()
}