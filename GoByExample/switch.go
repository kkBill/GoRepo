package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 2
	fmt.Print("write ", i , " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("today is " + time.Now().Weekday().String())
	fmt.Println("time now is " + time.Now().String())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatKindOfType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("it's type %s\n", t)
		}
	}
	whatKindOfType(12)
	whatKindOfType(true)
	whatKindOfType("String")
}