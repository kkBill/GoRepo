package main

import (
	"errors"
	"fmt"
)

/**
在go中，通常用显式的、独立的返回值来表示错误信息。
有点像C语言中的error number。但和Java中的异常不太一样
一般来讲，errors是最后一个返回值，其类型为 error ，这是一个内置的接口
*/

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New(msg) 创建一个基本的error value，错误信息在msg中给出
		return -1, errors.New("can't work with 42")
	}
	// 如果没有错误，则error返回nil
	return arg, nil
}

// 除了像上面这样定义错误，还可以自定义errors类型（需要实现Error()方法）
// 如下：
type customError struct {
	arg int
	msg string
}

// implement Error() on customError, Error() 是内置的接口
func (ce customError) Error() string {
	return fmt.Sprintf("%d - %s", ce.arg, ce.msg)
}

func f2(arg int) (int, error) {
	if arg == 42{
		// 返回自定义的 error
		return -1, customError{arg:arg, msg:"can't work with it"}
	}
	return arg, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	// 这里的 e.(customError) 是 类型断言（type assertion）
	// 如果想要获取自定义error中的数据（即取出customError类中的arg和msg）
	// 则需要通过 type assertion 的方式来获取
	// 关于 type assertion，可参考这里：https://books.studygolang.com/gopl-zh/ch7/ch7-10.html
	if ae, ok := e.(customError); ok{
		fmt.Println(ae.arg, ae.msg)
	}
}