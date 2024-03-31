package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个待取消的context
	ctx, cancel := context.WithCancel(context.Background())

	// 创建一个函数在新的协程中运行，模拟数据库查询操作
	go func() {
		// 假设数据库查询需要5秒才能完成
		time.Sleep(5 * time.Second)

		select {
		case <-ctx.Done():
			fmt.Println("main: ", ctx.Err())
			return
		default:
			fmt.Println("main: database query completed")
		}
	}()

	// 添加一个请求id到context中
	ctx = context.WithValue(ctx, "requestID", "123456")

	go ProcessRequest(ctx)

	// 在2秒后取消操作
	time.Sleep(10 * time.Second)
	cancel() //调用取消函数，将context标记为已取消
}

func ProcessRequest(ctx context.Context) {
	reqID, ok := ctx.Value("requestID").(string)
	if !ok {
		fmt.Println("No request ID found")
		return
	}

	fmt.Println("process request with ID:", reqID)

	// 模拟处理请求时，周期性检测context是否取消
	for {
		select {
		case <-ctx.Done():
			fmt.Println("request", reqID, "has been cancelled:", ctx.Err())
			return
		default:
			// 如果请求未被取消则继续运行
			fmt.Println("processing request", reqID)
			time.Sleep(1 * time.Second)
		}
	}
}
