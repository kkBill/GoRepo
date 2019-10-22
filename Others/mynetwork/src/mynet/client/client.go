package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil{
		log.Fatal(err)
	}
	log.Println("connect successfully...")
	defer conn.Close()

	//client := connector.CreateConnector(conn)

	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)

	//var data connector.DataBlock
	//data = connector.DataBlock{Name:"zju", Age:122}
	//err = client.Send(data)
	//if err!=nil{
	//	log.Printf("client send data err: %v\n",err)
	//}
	//
	//dataBlock, err := client.Receive()
	//if err!=nil{
	//	log.Printf("client receive data err: %v\n",err)
	//}
	//fmt.Printf("client: %s %d\n",dataBlock.Name, dataBlock.Age)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}