package main

import (
	"fmt"
	"reflect"
)

func demo1() {
	var f float64 = 3.14
	v := reflect.ValueOf(f)
	fmt.Println("type:",v.Type())
	fmt.Println("value:",v.Float())
	//fmt.Println(v.Elem()) // error
}

func demo2()  {
	// 函数 reflect.TypeOf 接受任意的 interface{} 类型, 并以reflect.Type形式返回其动态类型
	t := reflect.TypeOf(3)
	fmt.Println(t) // "int"
	fmt.Println(t.String()) // "int"

	//fmt.Printf("%T\n",3)
	//fmt.Printf("%T\n","hello")

	// 一个 reflect.Value 可以装载任意类型的值.
	// 函数 reflect.ValueOf 接受任意的 interface{} 类型, 并返回一个装载着其动态值的 reflect.Value.
	v := reflect.ValueOf(3)
	fmt.Println(v) // "3"
	fmt.Println(v.String()) // "<int Value>". 除非 Value 持有的是字符串, 否则 String 方法只返回其类型（对应下面语句2）
	fmt.Printf("%v\n",v) // "3"

	v = reflect.ValueOf("hello")
	fmt.Println(v) // "hello"
	fmt.Println(v.String()) // "hello" 语句2
	fmt.Printf("%v\n",v) // "hello"
}

func demo3() {
	//
	swap := func(in []reflect.Value) []reflect.Value{
		return []reflect.Value{in[1], in[0]}
	}

	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
		fn := reflect.ValueOf(fptr).Elem()
		// Make a function of the right type.
		v := reflect.MakeFunc(fn.Type(), swap)
		fn.Set(v)
	}

	// 函数声明
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0,1))
}

func demo4()  {
	add := func(in []reflect.Value) []reflect.Value {
		result := in[0].Interface().(int) + in[1].Interface().(int)
		return []reflect.Value{reflect.ValueOf(result)}
	}
	var addPtr func(int, int) int
	fn := reflect.ValueOf(&addPtr).Elem()
	v := reflect.MakeFunc(fn.Type(), add)
	fn.Set(v)

	fmt.Println(addPtr(1,2))
}

func main()  {
	//demo1()
	//demo2()
	demo4()
}
