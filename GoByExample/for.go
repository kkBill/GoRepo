package main

import "fmt"

// for 是go中唯一的循环结构
func main(){
	i:=1
	for i<=3 {
		fmt.Println(i)
		i++
	}

	for j:= 4; j<=5; j++{
		fmt.Println(j)
	}

	// 相当于 while(true)
	count := 0
	for {
		count++
		fmt.Println("xxx")
		if count == 5{
			break
		}
	}
}
