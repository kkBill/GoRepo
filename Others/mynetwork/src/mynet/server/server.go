package main

import (
	"fmt"
	"log"
	"mynet/connector"
	"net"
)

func main() {
	// 服务器监听的ip+port
	addr := "localhost:2000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() // 稍后回收

	for {
		// 等待连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		// 在一个新的 goroutine 中处理响应业务
		// 这样，外面这一个 goroutine 就可以继续监听了，从而达到并发处理的效果
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	// 业务逻辑
	fmt.Println("开始处理业务...")
	//io.Copy(c, c)
	server := connector.CreateConnector(c)
	dataBlock, err := server.Receive() // 卡住了
	fmt.Println("server after Receive()...")
	if err != nil {
		log.Printf("server receive data error: %v\n", err)
		c.Close()
		return
	}
	fmt.Printf("server: %s %d\n",dataBlock.Name, dataBlock.Age)
	dataBlock.Name = dataBlock.Name + " XXX"
	dataBlock.Age = dataBlock.Age * 2
	err = server.Send(dataBlock)
	fmt.Println("server after Send()...")
	if err != nil {
		log.Printf("server send data error: %v\n", err)
		c.Close()
		return
	}
	fmt.Println("业务处理完成...")
	// 处理完成后关闭连接
	c.Close()
}
