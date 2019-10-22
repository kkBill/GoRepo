package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name string
	age  int
}

// 实现fmt.Stringer
func (s Student) String() string {
	return "name:" + s.name + ", age:" + strconv.Itoa(s.age)
}

func main() {
	var v interface{}
	var i int = 23
	var s string = "hello, zju"
	v = i
	fmt.Println(v)
	v = s
	fmt.Println(v)

	Bob := Student{name:"Bob", age:24}
	fmt.Println(Bob)
}
