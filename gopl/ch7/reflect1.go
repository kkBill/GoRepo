package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type Foo interface {
	sayHi(s string)
}

type circle struct{
	radius float64
}

func (c circle) area() float64  {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) sayHi(s string) {
	fmt.Println("Hi, I am ", s)
}

func main() {
	c := circle{radius:3} // 初始化一个 circle 对象
	show(c)

	var f Foo
	f = c
	f.sayHi("circle")
}

func show(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}