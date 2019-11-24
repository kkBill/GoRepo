package main

import "fmt"

type Employee struct {
	ID            int
	Name, Address string
	Salary        int
}

type TreeNode struct {
	val         int
	left, right *TreeNode
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int // 辐条
}

func main() {
	bob := Employee{21951000, "Bob", "zhejiang", 15000}
	fmt.Println(bob)

	// p1 是一个Point结构体
	p1 := Point{X: 1, Y: 2}
	fmt.Println(p1) // {1,2}
	fmt.Printf("%#v\n",p1) //main.Point{X:1, Y:2} ，把各个字段也打印出来


	// pp 是一个指向Point{1,2}结构体的指针
	pp := &Point{X: 1, Y: 2}
	fmt.Println(pp)  // &{1,2}
	fmt.Println(*pp) // {1,2}
	fmt.Println(pp.Y)

	pp2 := new(Point)
	*pp2 = Point{1, 2}

	var w Wheel
	w.X = 8      //w.Circle.Point.X = 8
	w.Y = 8      //w.Circle.Point.Y = 8
	w.Radius = 5 //w.Circle.Radius = 5
	w.Spokes = 20

	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 6, Y: 6},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Println(w2)
	fmt.Printf("%#v\n", w2)
}
