package main

import "fmt"

type rect struct{
	width int
	height int
}

// 计算面积
func (r *rect) area() int  {
	return r.height * r.width
}

// 计算周长
func (r rect) perimeter() int {
	return (r.height + r.width) * 2
}

func main()  {
	r := rect{width:10, height:5}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
}