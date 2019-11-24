package main

import "fmt"

func main() {
	score := make(map[string]int)
	score["math"] = 98
	score["cs"] = 90
	score["cs"] = 78 // 覆盖前一条记录，key是不允许重复出现的

	fmt.Println(score)

	ages := map[string]int{
		"kkbill": 24,
		"Cindy": 18, // 不要忘记","
		"Bob": 35,
		"James": 35,
	}
	for name, age := range ages{
		fmt.Println(name, age)
	}


	mp := map[string]int{}
	mp["aaa"] = 233
	fmt.Println(mp["bbb"])
	delete(mp, "ccc")
}
