package main

import (  
    "runtime"
    "time"
)

var _ = runtime.GOMAXPROCS(3)

var a, b int

func u1() {  
    a = 1
	time.Sleep(3*1e9)
    b = 2
}

func u2() {  
    a = 3
    b = 4
}

func p() {  
    println(a)
    println(b)
}

func main() {  
    go u1()
    go u2()
    go p()
    time.Sleep(1 * time.Second)
    // 多次执行可显示以下以几种打印结果
    // 1   2
    // 3   4
    // 0   2 (奇怪吗？)
    // 0   0    
    // 1   4 (奇怪吗？)
}