package znet

import (
	"ZinxDemo01/Zinx/Ziface"
	"net"
)

// Connection 当前连接的模块
type Connection struct {
	// 当前连接的Socket TCP 套接字
	Conn *net.TCPConn

	// 当前连接的ID
	ConnID uint32

	// 当前连接的状态
	isClosed bool

	// 当前连接的绑定的处理业务的方法
	handleAPI Ziface.HandleFunc

	// 告知当前连接退出的Channel
	ExitChan chan bool
}

// NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI Ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callbackAPI,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
