package main

import (
	"fmt"
)

/**
range在不同数据类型上的作用
用于迭代iteration
*/
func main() {
	nums := []int{11, 22, 33}
	sum := 0
	// range on arrays and slices provides
	// both the index and value for each entry.
	for index, ele := range nums {
		fmt.Printf("[%d] = %d\n", index, ele)
		sum += ele
	}
	fmt.Println("sum:", sum)

	// range on map iterates over k/v pairs
	kvs := map[string]int{"aaa": 1, "bbb": 2, "ccc": 3}
	for k, v := range kvs {
		fmt.Printf("[%s] = %d\n", k, v)
	}
	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//range on strings iterates over Unicode code points.
	// 第一个返回值是字符对应的(首个字节的)下标
	// 第二个返回值是字符本身对应的ASCII码
	// 'a' --> 97
	for i, c := range "浙江大学" { // 每个中文字符占 3 byte
		fmt.Println(i, c) // 因此下标i: 0,3,6,9
	}
	for i, c := range "zju" {
		fmt.Println(i, c)
	}

	//fmt.Printf("%d",'a');
}
