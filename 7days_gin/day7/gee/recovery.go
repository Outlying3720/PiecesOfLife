package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)


func Recovery() HandleFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")

	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		str.WriteString(fmt.Sprintf("\n\t%s:%d", frame.File, frame.Line))
		if !more {
			break
		}
	}

	return str.String()
}