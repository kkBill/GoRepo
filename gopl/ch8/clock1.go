package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// 死循环，服务器端始终保持监听
	// 但此时服务器程序同一时间只能处理一个客户连接
	// 若要支持并发，则需要在 handleConn 处增加 go 关键字，
	// 让每一个对请求的处理都在新的goroutine中独立进行
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 处理请求
		// handleConn(conn)
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		// 由于net.Conn实现了io.Writer接口，我们可以直接向其写入内容。
		// 注意这里的时间格式必须这样写！（回看gopl）
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
