package main

import (
	"fmt"
	"time"
)

/**
We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.
When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
我们可以使用 channel 来作为 线程之间的同步工具
在这个例子中，我们使用 channel 的 blocking receive 来等待另一个线程执行结束
如果要等待多个线程都执行完，则最好使用 WaitGroup
*/

// 这个函数在另一个 goroutine 中执行，当执行完后，
// 通过channel 发送数据，用来通知其他 goroutine ——"这个线程已经执行好啦"
func worker(c chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	c <- true // send a value to notify that we are done
}

func main() {
	c := make(chan bool, 1)
	go worker(c)

	// 这里会阻塞，直到接收到来自channel c中的消息
	// 如果移除了下面这句话，程序会在worker()还没执行完就已经结束
	<-c
}
