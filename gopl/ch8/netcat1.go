package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	// 从conn中读取数据，并写进标准输出中
	mustCopy(os.Stdout, conn)
}

// 从连接中读取数据，并将读到的内容写到标准输出中，
// 直到遇到end of file的条件或者发生错误。
// 由于net.Conn实现了io.Writer/io.Reader接口，我们可以直接向其写入或读取内容。
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}