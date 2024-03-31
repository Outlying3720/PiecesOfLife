package main

import (
	"flag"
	"fmt"
    "net/http"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 && flag.NArg() != 3 {
		panic("usage: host port [path]")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	filePath := "."
	if path:=flag.Arg(2); len(path)>0 {
		filePath = path
	}
	fmt.Println("Start at: ", hostAndPort, " List file on: ", filePath)
	http.ListenAndServe(hostAndPort, http.FileServer(http.Dir(filePath)))
}
