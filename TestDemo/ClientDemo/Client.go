package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/
func main() {
	fmt.Println("Client Start...")

	// 创建TCP连接,得到Conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("net.Dial Error : ", err)
		return
	}

	for {
		// 调用Write写数据
		_, err = conn.Write([]byte("Hello,Zinx Server!"))
		if err != nil {
			fmt.Println("conn.Write Error : ", err)
			return
		}

		buf := make([]byte, 512)
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read Error : ", err)
			return
		}
		fmt.Printf("Server Back: %s\n", buf[:read])

		time.Sleep(2 * time.Second)
	}

}
