package main

import "fmt"

/**
当使用channel作为函数参数时，可以指定这个channel是用于发送数据的，还是接收数据的
通过这种方式可以增加程序的安全性
*/

// ping()函数只用来向channel中发送数据
func ping(sendChan chan<- string, msg string) {
	sendChan <- msg

	// 正是因为在函数参数中明确指定了该channel 为chan <- string（关键是方向）
	// 即表示它只能用来 发送数据
	// 因此下面这个语句编译就无法通过，而不需要等到运行时
	// 会提示 send-only type chan <- string
	//msg = <-sendChan
}

func pong(recvChan <-chan string, sendChan chan<- string) {
	msg := <-recvChan
	sendChan <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
