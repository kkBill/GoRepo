package main

import (
	"fmt"
	"sort"
)

// go 语言提供了 sort()函数

func main() {
	strs := []string{"a", "c", "b"}

	// 需要注意的是：sort()方法是 in-place 的
	// 也就是说，不会产生一个新的 strs，而是对原strs进行处理并返回
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{2,5,1,3}
	sort.Ints(ints)
	fmt.Println(ints)

	// 还可以利用sort来判断一个序列是否有序
	flag := sort.IntsAreSorted(ints)
	fmt.Println(flag)
}