package main

import (
	"fmt"
)

func main() {
	// make a slice of strings of length 3(initially zero-valued)
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "zju"
	s[1] = "zjnu"
	s[2] = "pku"
	fmt.Println("s:", s)
	fmt.Println("len:", len(s)) // slice的长度
	fmt.Println("s[2]:", s[2])
	//fmt.Println(s[3])

	// slice 中的 append 函数 可以在slice的末尾添加元素，并返回新的slice
	s = append(s, "thu")
	s = append(s, "cmu", "xxx")
	fmt.Println("after append:", s)

	// copy 函数可以把一个 slice 复制到另一个 slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)

	// slice支持以下语法，这个和python很像
	// slice[low:high]，遵守前闭后开原则
	l := s[2:5] // s[2], s[3], s[4]
	fmt.Println("s[2:5]:", l)

	l = s[:5] // s[0~4]
	fmt.Println("s[:5]:", l)

	l = s[2:] // s[2~ len(s)-1]
	fmt.Println("s[2:]:", l)

	// 声明的时候就进行初始化
	t := []string{"aaa", "bbb", "ccc", "ddd"}
	fmt.Println("t:", t)
	fmt.Println("t[1:4]:", t[1:4])

	// 利用slice来声明二维数组，3 指的是共有3个一维的slice
	// 至于每个一维的slice有多长，是可变的
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		matrix[i] = make([]int, innerLen) // innerLen 为一维长度
		for j := 0; j < innerLen; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("matrix:",matrix)
}
