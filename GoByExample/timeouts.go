package main

import (
	"fmt"
	"time"
)

func main() {
	// 假设我们正在执行一个外部调用，这个外部调用耗时2秒钟
	// 主线程需要用到该外部调用的结果，子线程通过channel把结果传回主线程
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// select 会执行第一个就绪的时间
	// 在下面的例子中，由于c1中传回结果需要2秒，
	// 而time.After()（该函数返回一个 <- chan Time）设置的超时时间为1秒
	// 因此，select会优先执行 超时事件
	// 如果把超时时间设为 3s，则会打印res
	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(1*time.Second):
		fmt.Println("timeout 1s")
	}


}
