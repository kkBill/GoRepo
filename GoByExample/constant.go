package main

import (
	"fmt"
	"math"
)

// go支持的常量类型有 字符、字符串、布尔、数值

const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 100

	const d = 233 / n
	fmt.Println(d)

	fmt.Println(int64(n))
	fmt.Println(math.Sin(n))
}
