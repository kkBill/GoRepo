package main

import (
	"fmt"
	"time"
)

/**
Timers are for when you want to do something once in the future
Tickers are for when you want to do something repeatedly at regular intervals.
*/

func main() {
	// 创建一个 ticker，每隔500毫秒触发一次
	// Tickers use a similar mechanism to timers:
	// a channel that is sent values.
	ticker := time.NewTicker(500 * time.Millisecond) // ticker 的时间是 <- chan Time

	done := make(chan bool)
	c := make(chan bool)

	// 子线程中的 <-ticker.C 每隔 500ms 响应一次，
	// 由于在主线程中设置了：在开始 1600ms后关闭ticker
	// 因此 子线程中的 ticker.C 只会响应3次
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done!")
				c <- true
				return
			case t := <-ticker.C: // t 的类型是 Time
				fmt.Println("tick at:", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	<- c
	fmt.Println("ticker stop!")
}