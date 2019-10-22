package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("emp:", arr)

	arr[4] = 100
	fmt.Println("set:", arr)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare and initialize:", b)

	matrix := [3][4]int{{1,2,3,4},{4},{6,7,8,9}}
	fmt.Println("matrix:", matrix)
}
