package main

import "fmt"

func main() {
	// we’ll use a jobs channel to communicate work to be done from
	// the main() goroutine to a worker goroutine.
	// When we have no more jobs for the worker
	// we’ll close the jobs channel.
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here’s the worker goroutine.
	// It repeatedly receives from jobs with j, more := <-jobs.
	// In this special 2-value form of receive,
	// the more value will be false if jobs has been closed
	// and all values in the channel have already been received.
	// We use this to notify on done when we’ve worked all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more{
				fmt.Println("received job:",j)
			}else{
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// in main() goroutine,
	// sends 3 jobs to the worker over the jobs channel,
	// then closes it.
	for i:=0; i<3;i++{
		jobs <- i
		fmt.Println("sent job:",i)
	}
	close(jobs)// 表示没有数据可发送了
	fmt.Println("sent all jobs")

	// 阻塞，等待worker goroutine全部完成后才结束
	<- done
}
