package znet

import (
	"ZinxDemo01/Zinx/Ziface"
	"fmt"
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

// StartReader 连接读的业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is Running....")
	defer fmt.Printf("ConnID: %d Reader is Exit, Remote Addr is : %s", c.ConnID, c.RemoteAddr().String())
	defer c.Stop()

	for {
		// 建立阻塞读取客户端数据到buf中
		buf := make([]byte, 512)

		read, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("c.Conn.Read Error: %s\n", err)
			continue
		}

		// 调用当前连接绑定的HandleAPI
		if err := c.handleAPI(c.Conn, buf, read); err != nil {
			fmt.Printf("c.ConnID: %d , Handle Error: %s\n", c.ConnID, err)
			break
		}

	}

}

// Start 启动连接 让当前连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Connection START...")

	// 启动当前连接读数据的业务
	go c.StartReader()
}

// Stop 停止连接 结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Printf("Connection STOP.... , ConnID: %d\n", c.ConnID)

	//如果当前连接已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true

	c.Conn.Close()
	close(c.ExitChan)
}

// GetTCPConnection GetTCPConnetion 获取当前链接绑定的 Socket Conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接模块的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端连接的TCP状态
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据 将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {

	return nil
}
