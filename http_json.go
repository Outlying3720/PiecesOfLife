package main

import (  
    "fmt"
    "net/http"
    "io/ioutil"
	"encoding/json"
)

type Ipdata struct {
	Ip string
}

func main() {  
    resp, err := http.Get("https://api.ipify.org?format=json")

    if resp != nil {
        defer resp.Body.Close()     // ok，即使不读取Body中的数据，即使Body是空的，也要调用close方法
    }

    //defer resp.Body.Close()       // （1）Error:在nil值判断之前使用，resp为nil时defer中的语句执行会引发空引用的panic
    if err != nil {
        fmt.Println(err)
        return
    }

    //defer resp.Body.Close()       // （2）Error:排除了nil隐患，但是出现重定向错误时，仍然需要调用close
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

	ipdata := Ipdata{}
	json.Unmarshal(body, &ipdata)

    fmt.Println(ipdata.Ip)
}