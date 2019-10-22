package main

import (
"fmt"
"math"
)
// 定义接口
type geometry interface {
	area() float64
	perimeter() float64
}

type rect2 struct {
	width, height float64
}

type circle struct{
	radius float64
}

// 对 rect2 类型实现接口
// 为了实现接口，就需要实现接口中定义的所有方法
func (r rect2) area() float64  {
	return r.height * r.width
}

func (r rect2) perimeter() float64 {
	return (r.height + r.width) * 2
}

// 对 circle 类型实现接口
func (c circle) area() float64  {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main()  {
	r := rect2{width:10, height:5}
	c := circle{radius:3}
	measure(r)
	measure(c)
}