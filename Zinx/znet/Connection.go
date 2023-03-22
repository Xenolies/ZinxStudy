package znet

import (
	"ZinxStudy/Zinx/utils"
	"ZinxStudy/Zinx/ziface"
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

	// 消息处理 MgsID 和对应的处理业务的关系
	MsgHandler ziface.IMsgHandler

	// 信息读写Channel 无缓冲
	msgChan chan []byte
}

// NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		MsgHandler: msgHandler,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
	}
	return c
}

// StartWriter 链接写的业务
func (c *Connection) StartWriter() {
	fmt.Println("Writer Goroutine is  running")
	defer fmt.Println(c.RemoteAddr().String(), " conn writer exit!")

	// 阻塞等待 Channel 消息
	for {
		select {
		case data := <-c.msgChan:
			// 如果有数据将数据写入客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Write Data Error: ", err)
				return
			}

		case <-c.ExitChan:
			// ExitChan 告知退出 Writer也要退出
			return

		}
	}

}

// StartReader 连接读的业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is  running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()
	for {
		// 创建拆包解包的对象
		dp := NewDataPack()

		//读取客户端的Msg head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("Read Msg Head Error: ", err)
			break
		}

		//拆包，得到msgid 和 datalen 放在msg中
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
			msg:  msg, //将之前的buf 改成 msg
		}

		// 判断工作池是否存在,存在则将消息交给工作池处理
		if utils.GlobalObject.WorkerPoolSize > 0 {
			// 开启工作处将消息交给工作池处理
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//从路由 Routers 中找到注册绑定Conn的对应Handle
			// 根据绑定好的MsgID, 传入消息处理模块 找到对应的处理业务
			go c.MsgHandler.DoMsgHandler(&req)

		}

	}
}

// Start 启动连接 让当前连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Connection START...")

	// 启动当前连接读数据的业务
	go c.StartReader()

	// 启动当前连接写数据的业务
	go c.StartWriter()
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

	// 告知Writer关闭
	c.ExitChan <- true

	// 将Channel关闭
	close(c.ExitChan)
	close(c.msgChan)
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
		return errors.New("msg Pack Error")
	}

	c.msgChan <- binaryMsg

	return nil
}
