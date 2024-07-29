package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

func writeLog(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "")
	buf := []byte(content)
	fd.Write(buf)
}

func main() {
	runtime.GOMAXPROCS(50)

	var w sync.WaitGroup
	count := int32(0)
	w.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				// count++

				// atomic.AddInt32(&count, 1)

				writeLog(".", "./cas2.log")
			}
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println(count)
}
