package main

import "fmt"

/**
Go 语言支持结构体，和C语言中的结构体是一样的，只是写法稍有不同
C/C++语言中：
struct Person{
	string name;
	int age;
}

Go中：
type Person struct{
	name string
	age int
}
 */

type Person struct{
	name string
	age int
}

// create a person and return a pointer of this object
func newPerson(n string, a int) *Person{
	p := Person{name:n, age:a}
	return &p
}

func main()  {
	// 创建一个结构体，可以对指定的域进行初始化，如果省略，则默认为0值
	fmt.Println(Person{"Cindy",24}) //{Cindy 24}
	fmt.Println(Person{name:"Bob", age:23}) //{Bob 23}
	fmt.Println(Person{name:"Jack"}) // {Jack 0}
	fmt.Println(newPerson("Curry",12)) // &{Curry 12}

	person := Person{name:"Kobe",age:39}
	fmt.Println(&person) //注意，这里和C不同，不是输出person1结构体的地址，而是输出 &{Kobe 39}
	fmt.Println(person.name, person.age) // Kobe 39

	p := &person
	fmt.Println(p) // &{Kobe 39}
	fmt.Println(p.name, p.age) // Kobe 39

	p.name = "James"
	fmt.Println(*p) // {James 39}
	fmt.Println(person) // {James 39}
}
