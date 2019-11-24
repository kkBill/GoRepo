package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1 // 新数组的len
	if zlen <= cap(x) {
		// 容量足够，就在原slice上扩展
		z = x[:zlen] // 这里这么写是为了确保语句(1)不会发生panic
	} else {
		// 容量不够，需要进行扩容（默认两倍进行扩容）
		zcap := zlen
		if zcap < 2*cap(x) {
			zcap = 2 * cap(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // copy x to z
	}
	z[zlen-1] = y // 语句(1)
	return z
}

func main() {
	months := [...]string{1: "January ", 2: "February ", 3: "March",
		4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}

	fmt.Println(len(months), cap(months)) // 13, 13

	Q2 := months[4:7]             // [4,7)
	fmt.Println(Q2)               // [April May June]
	fmt.Println(len(Q2), cap(Q2)) // 3, 9

	months[5] = "MayII"
	fmt.Println(Q2) // [April MayII June]

	a := [...]int{1, 2, 3, 4, 5}              // 声明并初始化一个数组
	fmt.Printf("%T %T %T\n", a, a[:], a[2:4]) // [5]int, []int, []int
	reverse(a[:])
	fmt.Println(a)

	s := []int{11, 22, 33, 44, 55} // 声明并初始化一个slice
	reverse(s)
	fmt.Println(s)
	fmt.Println(s[:len(s)]) // [0, len(s) )
	//f1 := 3.14
	//i1 := int(f1) // 类型转换
	//fmt.Println(f1, i1)

	// make函数创建一个指定元素类型、长度和容量的slice
	//ss1 := make([]string, 10)
	//ss2 := make([]string, 10 ,15)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d \t%v\n", i, cap(y), y)
		x = y
	}

}
