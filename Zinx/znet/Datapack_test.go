package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// DataPack 解包封包功单元测试
func TestNewDataPack(t *testing.T) {
	/*
		模拟服务器
	*/
	//  创建 SocketTCP
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(" net.Listen Error: ", err)
		return
	}
	// 创建Go承载负责从客户端处理业务
	go func() {
		// 从客户端读数据拆包处理
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(" listener.Accept() Error: ", err)
				return
			}

			// 处理客户端请求
			go func(conn net.Conn) {
				// 处理客户端请求
				// ----> 拆包流程<----//
				// 1 第一次从conn读,将head读出来
				// 2 第二次从conn读,根据head中的DataLen读data

				for {
					dp := NewDataPack() // 拆包对象 dp
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("Server io.ReadFull Error: ", err)
						// 一旦没有数据跳出
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("dp.Unpack Error: ", err)
						return
					}

					if msgHead.GetMsgLen() > 0 {
						// 说明 Message中有数据 需要进行第二次读取
						msg := msgHead.(*Message) // 类型断言 将msgHead转为 Message 类型
						msg.Data = make([]byte, msg.GetMsgLen())

						// 根据DataLen的长度开再从io中读取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("Message Unpack Data Error: ", err)
							return
						}

						// 完成的消息读取完毕
						fmt.Println("---> MsgID: ", msg.ID, ",DataLen: ", msg.DataLen, ",Data: ", msg.Data)
					}

				}
			}(conn)
		}
	}()

	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("net.Dial Error: ", err)
		return
	}

	// 创建一个封包对象 dp
	dp := NewDataPack()

	/*
		模拟粘包,封装两个Msg一起发送
	*/
	// 封装第一个包
	msg1 := &Message{
		ID:      1,
		DataLen: 5,
		Data:    []byte("12345"),
	}
	sendData1, err := dp.Pack(msg1)
	fmt.Println("sendData1: ", sendData1)
	if err != nil {
		fmt.Println(" dp.Pack msg1 Error: ", err)
		return
	}

	// 封装第二个包
	msg2 := &Message{
		ID:      2,
		DataLen: 7,
		Data:    []byte("1234567"),
	}
	sendData2, err := dp.Pack(msg2)
	fmt.Println("sendData2: ", sendData2)
	if err != nil {
		fmt.Println(" dp.Pack msg2 Error: ", err)
		return
	}

	// 将两个包站在一起
	sendData1 = append(sendData1, sendData2...)
	fmt.Println(sendData1)

	// 一同发送至服务端
	_, err = conn.Write(sendData1)
	if err != nil {
		fmt.Println("Client conn.Write Error: ", err)
		return
	}

	// 客户端阻塞
	select {}
}
