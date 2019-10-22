package main

import "fmt"

// go语言支持函数返回多个值
func multiReturnVals()(int, float32)  {
	return 1024, 3.14
}
func main()  {
	a, b := multiReturnVals()
	fmt.Println(a,b)

	// _ is blank identifier
	_, c := multiReturnVals()
	fmt.Println(c)
}
