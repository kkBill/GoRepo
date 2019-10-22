package main

import (
	"fmt"
	"time"
)

/**
定时器 timer
*/

func main() {
	// timer represent a single event in the future
	// you can tell the timer how long you want to wait
	// and it provide a channel that will be notified at that time
	timer1 := time.NewTimer(2 * time.Second)

	// the timer1.C blocks on the timer's channel C
	// until it sends a value indicating that timer expired
	<-timer1.C
	fmt.Println("timer 1 expired")

	// 除了向上面那样正常执行
	// 还可以提前把timer终止掉
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()
	// 我们让timer2 在另一个线程等待，但是在主线程中又提前把它终止掉了
	//ime.Sleep(2 * time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
}
