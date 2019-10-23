package main

import (
	"fmt"
	"reflect"
)

func demoKind() {
	// 空接口（interface{}）可以存放任意的值
	for _, v := range []interface{}{"hi", 23, func() {}} {
		v := reflect.ValueOf(v)

		// 通过Kind()方法确定存放在interface{}中的值的类型
		// 然后再根据特定的类型调用特定的方法取出相应的结果
		// 如果v调用的方法与v本身的类型不匹配（比如v是reflect.String，但调用了v.Int()）
		// 就会造成运行时异常（run-time panic）
		switch v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
			// fmt.Println(v.Int()) // cause run-time panic
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind: %s\n", v.Kind())
		}
	}
}

func demoMakeFunc() {
	// swap 这个函数实现是要传给 MakeFunc的
	// 它的传入传出参数必须是 reflect.Value，
	// 因为在调用这个方法前，不知道会传入什么具体的参数类型
	swap := func(inArgs []reflect.Value) []reflect.Value {
		return []reflect.Value{inArgs[1], inArgs[0]}
	}

	// makeSwap 将指针 fptr 设置为使用MakeFunc创建的新函数。
	// 当这个函数被调用的时候，反射 会将参数转化成reflect.Value，并调用 swap
	// 然后将swap函数的结果切片转换成由新函数返回的值
	makeSwap := func(fptr interface{}) {
		// fptr 是一个函数指针
		// 获取函数本身的值作为reflect.Value，然后查询其类型并设置值
		fn := reflect.ValueOf(fptr).Elem()
		fmt.Println(fn.Kind())  // func
		// fmt.Println(fn.Type())

		// fn.Type() 打印出来是 func(int, int) (int, int)
		// 也就是说，基于fn.Type()提供的函数参数类型，以及swap提供的函数体的实现，来创建一个函数
		v := reflect.MakeFunc(fn.Type(), swap)
		//fmt.Println(v.Kind()) // v Value类型的变量，它的Kind()是 func

		// ??
		fn.Set(v)
	}

	// 声明一个处理int型参数的swap函数
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(1,2))
}

func main() {
	//demoKind()
	demoMakeFunc()
}
