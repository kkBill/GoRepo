package main

import "fmt"
/**
Go 支持匿名函数，匿名函数可以形成闭包（closure）
 */

// intSeq()函数的返回值是一个函数，这个函数没有函数名
// 是一个匿名函数，在intSeq()函数体中定义
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main()  {
	// 当调用intSeq()函数时，其返回值也是一个函数
	// 即 nextInt 是函数名，而不是变量名，调用的时候要加括号
	nextInt := intSeq()

	// 为什么会造成这种情况的原理目前还不懂2019/09/30
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	anotherNextInt := intSeq()
	fmt.Println(anotherNextInt()) // 1
}
