package main

import "fmt"

// 比较两个slice类型是否相等
func equals(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	var a [3]int; // 默认情况下，数组的每个元素被初始化为0值
	fmt.Println(a[0])
	fmt.Println(len(a)) // len()为built-in函数

	var q [3]int = [3]int{111, 222, 333}
	a = q
	for i, v := range q { // 返回下标和对应的值
		fmt.Printf("%d %d\n", i, v)
	}

	b := [...]int{1, 2, 3}
	fmt.Printf("%T\n", b) //[3]int

	//q1 := [2]int{1,2}
	//q1 = [3]int{1,2,3}

	c := [...]string{0: "zju", 3: "pku", 1: "cmu"}
	for i, v := range c {
		fmt.Println(i, v)
	}

	d := [2]int{1, 2}
	e := [...]int{1, 2}
	f := [2]int{1, 3}
	fmt.Println(d == e, d == f, e == f) // true false false
}
