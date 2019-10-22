package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn2(conn)
	}
}

func echo(c net.Conn, msg string, delay time.Duration) {
	fmt.Fprintln(c, "\t",strings.ToUpper(msg))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", msg)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(msg))
}

func handleConn2(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		echo(c, scanner.Text(), 1*time.Second)
	}
	c.Close()
}
