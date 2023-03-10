package Ziface

import "net"

// Connection 定义连接模块接口
type Connection interface {
	// Start 启动连接 让当前连接准备开始工作
	Start()

	// Stop 停止连接 结束当前连接的工作
	Stop()

	// GetTCPConnetion 获取当前链接绑定的 Socket Conn
	GetTCPConnetion() *net.TCPConn

	// GetConnID 获取当前连接模块的ID
	GetConnID() uint32

	// RemoteAddr 获取远程客户端连接的TCP状态
	RemoteAddr() net.Addr

	// Sent 发送数据 将数据发送给远程的客户端
	Sent(data []byte) error
}

// HandleFunc 定义一个处理业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
