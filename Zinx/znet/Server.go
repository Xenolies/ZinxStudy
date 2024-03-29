package znet

import (
	"ZinxStudy/Zinx/utils"
	"ZinxStudy/Zinx/ziface"
	"fmt"
	"net"
)

// Server 实现 IServer 接口 ,定义一个服务器模块
type Server struct {
	// 服务器名称
	ServerName string
	// 服务器IP版本
	IpVersion string
	// 服务器端口
	IP string
	// 服务器监听端口
	Port int
	// 当前 Server 的消息处理模块, 用来绑定 MsgID 和对应的处理业务API
	MsgHandler ziface.IMsgHandler

	// 该Server的链接控制器
	ConnManager ziface.IConnectionManager

	// Server创建链接自动调用的 Hook函数
	OnConnStart func(Conn ziface.IConnection)

	//	Server销毁链接自动调用的 Hook函数
	OnConnStop func(Conn ziface.IConnection)
}

func (s *Server) Start() {

	fmt.Printf("[START] Server Listener at IP: %s , Port %d is Starting\n", utils.GlobalObject.Host, utils.GlobalObject.TcpPort)

	fmt.Printf("[Zinx] Version %s, MaxConn:%d, MaxPackeetSize:%d \n", utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)

	go func() {
		// 开启工作池 和 消息队列
		s.MsgHandler.StartWorkPool()

		// 获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("net.ResolveIPAddr Error : ", err)
			return
		}
		// 监听服务器地址
		tcpListener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("net.ListenTCP Error : ", err)
			return
		}

		var conId uint32
		conId = 0

		fmt.Println("Start Zinx Server Success! [", utils.GlobalObject.Name, "] is Listening")

		// 阻塞等待客户端链接和处理客户端链接业务(读写)
		for {
			conn, err := tcpListener.AcceptTCP()
			if err != nil {
				fmt.Println("tcpListener.AcceptTCP Error : ", err)
				continue
			}

			// 客户端链接后的读写操作
			// 将处理新链接的任务方法和Conn绑定得到连接模块
			dealConn := NewConnection(s, conn, conId, s.MsgHandler)
			conId++

			// 建立链接前判断是否超过最大链接个数
			// 超过就关闭
			if s.ConnManager.Len() > utils.GlobalObject.MaxConn {
				fmt.Println("[MAX]Too Many Connection!!")
				fmt.Println("Now Conn: ", s.ConnManager.Len(), " MaxConn Setting: ", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			//启动连接任务处理
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	// 服务器终止
	fmt.Println("[STOP] Zinx Server[", utils.GlobalObject.Name, "] is STOP!")
	s.ConnManager.ClearConn()

}

func (s *Server) Serve() {
	// 启动服务
	s.Start()

	// 建立阻塞状态
	select {}
}

// NewServer 初始化 Server 模块
func NewServer() ziface.IServer {
	s := &Server{
		// 使用 utils.GlobalObject 替换
		ServerName:  utils.GlobalObject.Name,
		IpVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandler:  NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
	return s
}

// GetConnMgr 返回当前Server中的ConnManager
func (s *Server) GetConnMgr() ziface.IConnectionManager {
	return s.ConnManager
}

// AddRouter 想着MsgHandler添加路由
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	// 将Router 添加到 MsgHandler 中
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Router Add Success!!")
}

// SetOnConnStart 注册Server创建链接自动调用的 Hook函数
func (s *Server) SetOnConnStart(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop 注册Server销毁链接自动调用的 Hook函数
func (s *Server) SetOnConnStop(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart 调用Server创建链接自动调用的 Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("-----> Call OnConnStart()... ")
		s.OnConnStart(conn)
	}
}

// CallOnConnStop 调用Server销毁链接自动调用的 Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("-----> Call OnConnStop()... ")
		s.OnConnStop(conn)
	}
}
