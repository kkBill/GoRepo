package main

import (
	"fmt"
	"time"
)

// 这里说的 worker pools 的概念，相当于其他语言中的 线程池
// 暂且这么理解吧

// Here’s the worker, of which we’ll run several concurrent instances.
// These workers will receive work on the jobs channel and
// send the corresponding results on results.
// We’ll sleep a second per job to simulate an expensive task.

func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "start job", j)
		time.Sleep(time.Second) // 模拟任务处理消耗的时间
		fmt.Println("worker", id, "finish job", j)
		// 返回结果
		results <- j * 2
	}
}

// 感觉理解的不深！
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动 3 个worker，初始时会被堵塞，因为还没有向jobs中写入数据
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	//
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // 关闭 jobs channel 表明数据已经发送完了

	for r := 1; r <= 5; r++ {
		<-results
	}
}
