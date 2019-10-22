package main

import "fmt"

/*
Go 支持指针操作，和C/C++中的概念是一摸一样的
在C/C++中，int 型指针 int * p = null;
在Go中，int 型指针    var p *int = nil
*/

func fakeSwap(a, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func realSwap(p1 *int, p2 *int) {
	temp := *p1
	*p1 = *p2
	*p2 = temp
}

func main() {
	a := 11
	b := 22
	fmt.Printf("a = %d, b = %d\n", a, b)
	fakeSwap(a, b)
	fmt.Printf("a = %d, b = %d\n", a, b)
	realSwap(&a, &b);
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println(&a, &b)
}
