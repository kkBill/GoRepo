package main

import (
	"fmt"
	"time"
)

func f(msg string)  {
	for i:=0; i<10;i++{
		fmt.Println(msg, ":", i)
	}
}

func main()  {
	// 主线程，f 里面的循环顺序执行
	f("direct")

	// 下面两个线程则是并发执行
	// 子线程1
	go f("goroutine")
	// 子线程2
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 主线程sleep 100毫秒后退出
	// 当主线程退出后，其他线程也会被迫退出（而不管有没有执行完）
	time.Sleep(1000)
	fmt.Println("done")
}