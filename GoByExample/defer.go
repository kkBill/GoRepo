package main

import (
	"fmt"
	"os"
)

/**
defer 关键字通常用于清理（cleanup）工作
有点像Java中finally关键字所起的作用

假设现在我们要创建一个file，向其写入数据并关闭
我们看看如何用 defer 关键字来处理
 */

func main() {
	// 在创建好文件之后，紧随其后的是 defer 关键字
	// defer 语句表明：在main()方法的退出前，去执行 closeFile()
	// 所以在writeFile()完成后，就会执行closeFile()
	f := createFile("hello-go.txt")
	defer closeFile(f)
	writeFile(f)
}

// 创建一个文件并返回文件描述符
func createFile(filename string) *os.File {
	fmt.Println("creating")
	fd, err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	return fd
}

// 向文件中写入一句话
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "hello go")
}

// 关闭文件
func closeFile(f *os.File)  {
	fmt.Println("closing")
	err := f.Close()
	if err!=nil{
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}