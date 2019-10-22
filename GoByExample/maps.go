package main

import "fmt"

func main() {
	// 利用make方法创建一个空的map
	// make(map[key]value)
	m := make(map[string]int)

	// 设置key/value的值
	m["k1"] = 7
	m["k2"] = 13

	// 输出key-value键值对，appear in the form map[k:v k:v k:v]
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	// 提取键对应的值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 从map中返回某个key对应的value时，有一个可选的第二返回值
	// 这个返回值是一个bool类型，表明当前这个key是否存在
	_, present := m["k2"]
	fmt.Println("present or not:", present)
	_, exist := m["k3"]
	fmt.Println(exist) // false

	//内置的delete方法可以删除map中的键值对
	delete(m, "k1")
	fmt.Println("after delete map:", m)

	// 直接在声明的时候进行初始化
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
