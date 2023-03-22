package main

import (
	"ZinxStudy/Zinx/znet"
	"fmt"
	"io"
	"net"
	"time"
)

/*
模拟客户端 01
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
		// 创建一个封包对象
		dp := znet.NewDataPack()
		// 封包要发送的信息
		binaryMsg, err := dp.Pack(znet.NewMessage(1, []byte("Hello,Zinx For Client01")))
		if err != nil {
			fmt.Println("Msg Pack Error: ", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("binaryMsg Write Error: ", err)
			return
		}

		// 1 先读取流中的Head 得到 ID 和 DataLen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("Read binaryHead Error: ", err)
			break
		}
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("binaryHead UnPack Error: ", err)
			break
		}

		// 2 根据 DataLen 读取发送回来的消息
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msgHead.GetMsgLen())

			// 从 Conn中读取消息内容
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("Message Data Read Error: ", err)
				break
			}
			fmt.Println("---> [FROM SERVER]:  MsgID: ", msg.ID, ", DataLen: ", msg.DataLen, ",Data: ", string(msg.Data))
		}

		time.Sleep(2 * time.Second)
	}

}
