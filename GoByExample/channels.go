package main

import "fmt"

/**
channel 可以理解为管道，用于go语言中线程间通信
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine
and receive those values into another goroutine.
*/

func main()  {
	// 通过make(chan value-type)来创建一个channel
	// channel的类型由 chan 和 其传递的值的类型 共同决定
	// 也就是说，如果一个channel作为参数上，应该这样写：
	// func foo(c chan string){}
	message := make(chan string)

	// 使用 channel <- 的形式向一个channel中发送数据
	// 在这里，我们开启一个新的 goroutine，
	// 并在这个 goroutine 中向 message 这个channel中发送数据
	go func() {message <- "ping"}()

	// <- channel 表示从一个channel中接收数据
	msg := <- message // 语句1
	fmt.Println(msg)
	/**
	By default sends and receives block
	until both the sender and receiver are ready.
	This property allowed us to wait at the end of our program
	for the "ping" message without having to use any other synchronization.
	默认情况下，发送和接收都会同步的，比如上面的语句1，会阻塞的进行直到接收到数据
	这种特性避免了额外使用同步机制
	 */
}