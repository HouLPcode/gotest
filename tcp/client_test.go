package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Fatal(err)
	}

	clientHandle(conn)
}

func clientHandle(c net.Conn) {
	if c == nil {
		return
	}

	buf := make([]byte, 1024)
	for {
		len, _ := c.Read(buf)
		fmt.Println("receive data: ", string(buf[:len]))
		c.Write([]byte("client msg"))
		time.Sleep(time.Second)
	}
}
