package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100*time.Millisecond)
	const n = 45
	result := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, result)

}

func spinner(delay time.Duration) {
	for{
		for _, r := range `-\|/` {
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2{
		return x
	}
	return fib(x-1) + fib(x-2)
}
