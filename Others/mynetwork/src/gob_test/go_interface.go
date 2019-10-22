package gob_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

type Point struct {
	X, Y int
}
func (p Point) Hypotenuse() float64 {
	// Hypot returns Sqrt(p*p + q*q)
	return math.Hypot(float64(p.X), float64(p.Y))
}
type Pythagoras interface {
	Hypotenuse() float64
}

// 这个例子展示了如何 encode/decode 一个接口类型(interface{})的值
// 与其他常规的类型（比如结构体）最大的不同在于：
// 需要注册一个明确的实现该接口的类型
func GobInterface()  {
	// 我们必须要对encoder和decoder注册具体的类型，
	// 通常来说，decoder和encoder是在不同的机器上的。
	// 经过“注册”，解析引擎才能知道实现这一接口的具体类型是什么
	// （因为同一个接口可以由多种不通过的实现，不然怎么解析？）
	gob.Register(Point{})

	var buf bytes.Buffer //
	// 创建一个 encoder
	encoder := gob.NewEncoder(&buf)
	// 发送数据，这里故意写的繁琐一点，方便看清楚分别是什么类型
	var p Pythagoras // var p Point 也可以，但是前者更好
	p = Point{3, 4}
	err := encoder.Encode(&p)
	if err != nil{
		log.Fatal("Encode error:",err)
	}

	// 创建一个 decoder
	decoder := gob.NewDecoder(&buf)
	// 接收数据（数据会写进 q 中）
	var q Pythagoras
	err = decoder.Decode(&q)
	if err != nil{
		log.Fatal("Decode error:",err)
	}
	fmt.Println(q.Hypotenuse()) // 5
}
