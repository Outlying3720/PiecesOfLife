package main

//5、从panic中恢复
//recover()函数能够捕获或拦截panic，但必须在defer函数或语句中直接调用，否则无效。

import "fmt"

func doRecover() {  
    fmt.Println("recovered =>",recover()) //prints: recovered => <nil>
}

func main() {  
    defer func() {
        fmt.Println("recovered:",recover()) // ok
    }()

	defer func() {
        fmt.Println("recovered 2:",recover()) // ok
    }()

    defer func() {
        doRecover() //panic is not recovered
    }()

    // recover()   //doesn't do anything
    panic("not good")
    // recover()   //won't be executed :)

    fmt.Println("ok")
}