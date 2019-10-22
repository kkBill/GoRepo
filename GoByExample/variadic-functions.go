package main

import "fmt"

// Variadic Functions 就是函数的参数可以是变化的
// 最典型的就是 C语言中的printf()方法
// 在golang 中也是一样
func sum(nums ...int)  {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main()  {
	sum(1,2)
	sum(3,4,5)

	nums := []int{1,2,3,4,5}
	sum(nums...) // 这种调用方式是第一次见
}