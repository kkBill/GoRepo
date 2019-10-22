package main

import (
	"fmt"
	"sort"
)

// 自定义排序
// 根据string的长度进行排序

// 想要自定义实现sort()方法，则必须实现Len()/Less()/Swap() 接口
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
//type Interface myInterface {
//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)
//}

type byLength []string

func (s byLength) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp

	// demo 上是这么写的，什么意思啊？不懂
	// 这是 Go 中特有的写法
	//s[i], s[j] = s[j], s[i]
	// 还可以写成
	// s[i], s[j] = s[j], s[i]
}

func (s byLength) Len() int{
	return len(s)
}

func (s byLength) Less(i, j int) bool{
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach","watermelon","banana"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

