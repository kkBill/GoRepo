package main

import "fmt"

func main(){
	// use var to declare 1 or more varables
	var a = "zju"
	fmt.Println(a)
	// declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)
	
	var x, y = "this is a string", 3.14
	fmt.Println(x, y)

	var d = true
	fmt.Println(d)
	
	// varables declared without initialization are zero-valued
	var e int
	fmt.Println(e) // 0

	// the := syntax is more shorthand for declaring and initializing a varable
	f := "apple" // ==> var f string = "apple"
	fmt.Println(f)
}
