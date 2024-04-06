package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"geerpc"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func (f Foo) Mul(args Args, reply *int) error {
	*reply = args.Num1 * args.Num2
	return nil
}

type Bar int

func (f Bar) A(args Args, reply *int) error {
	*reply = args.Num1
	return nil
}

func (f Bar) B(args Args, reply *int) error {
	*reply = args.Num1 - args.Num2
	return nil
}

func startServer(addr chan string) {
	var foo Foo
	var bar Bar
	if err := geerpc.Register(&foo); err != nil {
		log.Fatal("register failed: ", err)
	}
	if err := geerpc.Register(&bar); err != nil {
		log.Fatal("register failed: ", err)
	}

	l, _ := net.Listen("tcp", ":9999")
	geerpc.HandleHTTP()
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	_ = http.Serve(l, nil)
}

func call(addr chan string) {
	client, _ := geerpc.DialHTTP("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call(context.Background(), "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d:", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go call(addr)
	startServer(addr)
}
