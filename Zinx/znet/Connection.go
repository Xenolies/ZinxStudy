package znet

import (
	"ZinxDemo01/Zinx/ziface"
	"errors"
	"fmt"
	"io"
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
	// 告知当前连接退出的Channel
	ExitChan chan bool

	// 当前连接的Router处理
	Router ziface.IRouter
}

// NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
	return c
}

// StartReader 连接读的业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is Running....")
	defer fmt.Printf("ConnID: %d Reader is Exit, Remote Addr is : %s", c.ConnID, c.RemoteAddr().String())

	defer c.Stop()

	for {
		// 创建一个拆包解包对象
		dp := NewDataPack()

		// 读取客户端 Msg Head 8字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("Read Msg Head ErrorL ", err)
			break
		}

		// 拆包得到 MsgID和msgDataLen 放到msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("UnPack Message Error: ", err)
			break
		}

		// 根据 DataLen 读取客户端发送的数据
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("Read Msg Data Error: ", err)
				break
			}
		}

		msg.SetData(data)

		// 得到当前Conn数据的Request的请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		// 预注册路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		//在路由中找到注册绑定的Conn的Router调用

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

// GetTCPConnection GetTCPConnection 获取当前链接绑定的 Socket Conn
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

// SendMsg 提供一个SendMsg方法,
// 将要发送给客户端的数据线进行封包,然后再发送.
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	// 先判断Conn是否关闭
	if c.isClosed == true {
		return errors.New("connection Closed When Send Msg")
	}

	// 将 data 进行封包
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMessage(msgID, data))
	if err != nil {
		fmt.Println("msg Pack Error", err)
		return errors.New("Msg Pack Error")
	}

	// 将数据发送到客户端
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("Write binaryMsg Error: ", err)
		return errors.New("Write binaryMsg Error")
	}

	return nil
}
