package main

import (
	"fmt"
	"sync"
	"time"
)

/**
等待多个线程执行结束
 */

// 我们会在多个线程中调用这个函数
// 【特别注意】：WaitGroup必须要通过指针的方式传递
func myWorker(id int, wg *sync.WaitGroup)  {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)

	// 当一个线程执行完成后，就通过wg.Done()告知 wg
	wg.Done() // 内部实现很简单，就是wd.Add(-1)
}

func main()  {
	var wg sync.WaitGroup

	// 启动多个线程
	for i:= 1;i<=5;i++{
		wg.Add(1) // 每启动一个线程
		go myWorker(i, &wg)
	}

	// wait()函数会一直阻塞，直到 WaitGroup 内部的counter 变为0
	// 表示所有等待的线程均执行完毕
	wg.Wait()
}