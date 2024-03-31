package main

import (  
    "fmt"
	"time"
    // "runtime"
)

func main() {  
    done := false
    go func(){
		time.Sleep(1*1e8)
        done = true
    }()

    for !done {
        // ... 
        //runtime.Gosched() // 让scheduler执行调度，让出执行时间片
    }
    fmt.Println("done!")
}