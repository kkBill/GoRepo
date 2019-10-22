package connector

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"log"
	"net"
)

type Connector struct {
	conn net.Conn
}

type DataBlock struct {
	Name string
	Age  int
}

func CreateConnector(c net.Conn) *Connector {
	return &Connector{conn:c}
}

// 发送结构化数据，内部先对数据进行编码，再发送字节流数据
func (c *Connector) Send(data DataBlock) error {
	// 先对数据进行编码
	byteArray, err := encode(data)
	// 设置头部
	byteBuf := make([]byte, len(byteArray)+4)                     // +4 是为了存放该条数据的数据长度
	// 这个函数不太懂
	binary.BigEndian.PutUint32(byteBuf[:4], uint32(len(byteBuf))) // 前4个字节存放数据长度
	copy(byteBuf[4:], byteArray)                                  // 拷贝有效数据
	// 向 conn 中写入数据
	_, err = c.conn.Write(byteBuf)
	if err != nil {
		log.Printf("write error: %v\n", err)
		return err
	}
	return nil
}

// 接收数据，接收到的数据是字节流数据，但内部会对其进行解码
func (c *Connector) Receive() (DataBlock, error) {
	// 1. 先读取头部
	header := make([]byte, 4)
	// func ReadFull(r Reader, buf []byte) (n int, err error)
	// ReadFull 从 r 精确地读取 len(buf) 字节数据填充进 buf
	_, err := io.ReadFull(c.conn, header)
	if err != nil {
		log.Printf("read header data error: %v\n", err)
		return DataBlock{}, err
	}

	// 将头部字节序列转化成数值
	dataLen := binary.BigEndian.Uint32(header)
	dataBuf := make([]byte, dataLen)
	_, err = io.ReadFull(c.conn, dataBuf)
	if err != nil {
		log.Printf("read real data error: %v\n", err)
		return DataBlock{}, err
	}
	// 2. 对字节流数据进行解码，解析成可读数据
	dataBlock, err := decode(dataBuf)
	return dataBlock, err
}

// 编码，返回编码后的字节序列和错误码
func encode(data DataBlock) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf) // 数据会写进 buf 中
	err := encoder.Encode(data)
	if err != nil {
		log.Printf("encode error: %v\n", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解码，返回解码后的数据结构体和错误码
func decode(b []byte) (DataBlock, error) {
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf) // 将从 buf 中读取数据

	var data DataBlock
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("decode error: %v\n", err)
		return DataBlock{}, err
	}
	return data, nil
}
