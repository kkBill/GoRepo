package main

import "fmt"

/**
在之前的练习中，对数组、slice和map都应用了range进行迭代
事实上，We can also use this syntax to iterate
over values received from a channel.
 */

func main() {
	queue := make(chan string, 3)
	queue <- "aaa"
	queue <- "bbb"
	queue <- "ccc"
	close(queue)

	// 对queue中的收到的内容进行迭代
	for elem := range queue{
		fmt.Println(elem)
	}

	// 这个例子也表明，我们可以关闭一个非空的channel
	// 在关闭的channel中，仍然存在数据，可以被接收
}