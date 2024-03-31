package main

import (  
    "fmt"
    "time"
)

type field struct {  
    name string
}

func (p *field) print() {  
    // fmt.Println(p.name)
	fmt.Printf("%p", p.name)
}

func main() {  
    data := []field{{"one"},{"two"},{"three"},{"4"},{"5"},{"6"},{"7"},{"8"},{"9"},{"10"},{"11"},{"12"},{"13"},{"14"},{"15"}}
    for _,v := range data {
        // 解决办法：添加如下语句
        // v := v
        go v.print()
    }
    time.Sleep(3 * time.Second)     //goroutines print: three, three, three

    data2 := []*field{{"one"}, {"two"}, {"three"}}  // 注意data2是指针数组
    for _, v := range data2 {
        go v.print()                // go执行是函数，函数执行之前，函数的接受对象已经传过来
    }
    time.Sleep(3 * time.Second)     //goroutines print: one, two, three
}