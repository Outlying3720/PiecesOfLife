package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
)

func main() {
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true

	proxy.OnRequest().DoFunc(
		func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
			log.Printf(r.URL.Path)
			r.URL.Path = "/-----http://" + r.URL.Host + r.URL.Path
			r.URL.Host = "dlive-paper.daisyorge.fun"
			r.Host = "dlive-paper.daisyorge.fun"
			return r,nil
		})
		
	log.Printf("Start simple go proxy!")
    log.Fatal(http.ListenAndServe(":8081", proxy))
}