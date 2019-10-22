package gob_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

type R struct {
	Y, W int
}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func GobBasic() {
	// 初始化 encoder 和 decoder
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf) // will write to buf
	decoder := gob.NewDecoder(&buf) // will read from buf

	// Encode (send) some values
	err := encoder.Encode(P{X: 3, Y: 4, Z: 5, Name: "hello"})
	if err != nil {
		log.Fatal("Encode error:",err)
	}

	// Decode (receive) and print the values
	//var q Q
	//err = decoder.Decode(&q)
	//if err != nil {
	//	log.Fatal("Decode error:",err)
	//}
	//
	//// 注意，不能写成 q.X，因为在接收方，定义的是 int型 指针
	//// *(q.X) 与 *q.Y 结果相同，但前者语义更加明确
	//fmt.Printf("%d %d %s\n", *(q.X), *q.Y, q.Name)

	//var p P
	//err = decoder.Decode(&p)
	//if err != nil {
	//	log.Fatal("Decode error:",err)
	//}
	//// 这里的接收方和传入方格式完全一致
	//fmt.Printf("%d %d %d %s\n", p.X, p.Y, p.Z, p.Name)

	var r R
	err = decoder.Decode(&r)
	if err != nil {
		log.Fatal("Decode error:",err)
	}

	//fmt.Printf("%d %d %d %s\n", r.X, r.Y, r.Z, r.Name)
	// 会输出如下：因为接收端是根据字段名称进行匹配的
	// r.X undefined (type R has no field or method X)
	// r.Z undefined (type R has no field or method Z)
	// r.Name undefined (type R has no field or method Name)
	fmt.Printf("%d\n", r.Y) // ok
}