package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        ctx := req.Context()
        fmt.Println("server: got request")
        select {
        case <-time.After(2 * time.Second):
            w.Write([]byte("request processed"))
        case <-ctx.Done():
            err := ctx.Err()
            fmt.Println("server:", err)
            internalError := http.StatusInternalServerError
            http.Error(w, err.Error(), internalError)
        }
    })

    http.ListenAndServe(":8000", nil)
}
