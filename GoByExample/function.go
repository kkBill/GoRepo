package main

import "fmt"

func add(a int, b int) int  {
	return a+b
}

// 如果函数有多个参数且类型一致，
// 则可以只在最后一个参数后面指明参数类型即可
func add3num(a, b, c int) int  {
	return a+b+c
}

func main()  {
	fmt.Println("1 + 2 = ", add(1,2))
	fmt.Println("1 + 2 + 3 = ", add3num(1,2, 3))
}