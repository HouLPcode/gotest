package main

import (
	"bytes"
	"fmt"
	"net"
)

//终端执行 telnet localhost 1234 进行测试
func main() {
	svr, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("server starting...")

	conn, err := svr.Accept()
	if err != nil {
		fmt.Println(err)
	}

	handle(conn)
}

func handle(c net.Conn) {
	if c == nil {
		return
	}

	i := uint(0)
	//for {
	//	fmt.Println("hello world",i)
	//	time.Sleep(time.Second)
	//	i++
	//}

	buf := make([]byte, 1024)

	for {
		c.Write([]byte("i am the server\n"))

		len, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("len is ", len)
		if bytes.Equal(buf[:len], []byte("quit")) {
			fmt.Println("connet close ", i)
			break
		}
		fmt.Println(string(buf[:len]), i)
		i++
	}

	c.Close()
}
