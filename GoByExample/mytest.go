package main

import "fmt"

func swap(a []int, i, j int){

	// 方式1（最常规的）
	//tmp := a[i]
	//a[i] = a[j]
	//a[j] = tmp

	// 方式2（尚不理解）
	//a[i], a[j] = a[j], a[i]

	// 方式3（尚不理解）
	//a[i], a[j] = a[j], a[i]
}

func main() {
	ints := []int{1,2,3}
	swap(ints, 0, 1)
	fmt.Println(ints)
}
