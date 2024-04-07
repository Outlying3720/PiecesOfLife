package main

import (
	"context"
	"geerpc/registry"
	"geerpc/xclient"
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

func (f Foo) Sleep(args Args, reply *int) error {
	time.Sleep(time.Second * time.Duration(args.Num1))
	*reply = args.Num1 + args.Num2
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

func foo(xc *xclient.XClient, ctx context.Context, typ, serviceMethod string, args *Args) {
	var reply int
	var err error
	switch typ {
	case "call":
		err = xc.Call(ctx, serviceMethod, args, &reply)
	case "broadcast":
		err = xc.Broadcast(ctx, serviceMethod, args, &reply)
	}
	if err != nil {
		log.Printf("%s %s err: %v", typ, serviceMethod, err)
	} else {
		log.Printf("%s %s success: %d + %d = %d", typ, serviceMethod, args.Num1, args.Num2, reply)
	}
}

func call(registry string) {
	d := xclient.NewGeeRegistryDiscovery(registry, 0)
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
	defer func() { _ = xc.Close() }()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			foo(xc, context.Background(), "call", "Foo.Sum", args)
		}(i)
	}
	wg.Wait()
}

func broadcast(register string) {
	d := xclient.NewGeeRegistryDiscovery(register, 0)
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
	defer func() { _ = xc.Close() }()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			foo(xc, context.Background(), "broadcast", "Foo.Sum", args)

			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
			foo(xc, ctx, "broadcast", "Foo.Sleep", args)
		}(i)
	}
	wg.Wait()
}

func startRegistry(wg *sync.WaitGroup) {
	l, _ := net.Listen("tcp", ":9999")
	registry.HandleHTTP()
	log.Println("start rpc registry on", l.Addr())
	wg.Done()
	_ = http.Serve(l, nil)
}

func startServer(registryAddr string, wg *sync.WaitGroup) {
	var foo Foo
	var bar Bar
	server := geerpc.NewServer()
	if err := server.Register(&foo); err != nil {
		log.Fatal("register failed: ", err)
	}
	if err := server.Register(&bar); err != nil {
		log.Fatal("register failed: ", err)
	}

	l, _ := net.Listen("tcp", ":0")
	registry.Heartbeat(registryAddr, "tcp@"+l.Addr().String(), 0)
	wg.Done()
	log.Println("start rpc server on", l.Addr())

	server.Accept(l)
}

func main() {
	log.SetFlags(0)
	registryAddr := "http://localhost:9999/_geerpc_/registry"
	var wg sync.WaitGroup
	wg.Add(1)
	go startRegistry(&wg)
	wg.Wait()
	wg.Add(2)
	go startServer(registryAddr, &wg)
	go startServer(registryAddr, &wg)
	wg.Wait()

	time.Sleep(time.Second)
	call(registryAddr)
	broadcast(registryAddr)

}
