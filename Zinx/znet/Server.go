package znet

import (
	"ZinxDemo01/Zinx/utils"
	"ZinxDemo01/Zinx/ziface"
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
	// 当前Sever添加一个Router,Server注册的链接处理业务
	Router ziface.IRouter
}

func (s *Server) Start() {

	fmt.Printf("[START GlobalObject] Server Listener at IP: %s , Port %d is Starting\n", utils.GlobalObject.Host, utils.GlobalObject.TcpPort)

	fmt.Printf("[START] Server Listener at IP: %s , Port %d is Starting\n", s.IP, s.Port)

	go func() {
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

		fmt.Println("Start Zinx Server Success! ", s.ServerName, " is Listening")

		// 阻塞等待客户端链接和处理客户端链接业务(读写)
		for {
			tcpConn, err := tcpListener.AcceptTCP()
			if err != nil {
				fmt.Println("tcpListener.AcceptTCP Error : ", err)
				continue
			}

			// 客户端链接后的读写操作

			// 将处理新链接的任务方法和Conn绑定得到连接模块
			dealConn := NewConnection(tcpConn, conId, s.Router)
			conId++

			//启动连接任务处理
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	// 服务器终止
}

func (s *Server) Serve() {
	// 启动服务
	s.Start()

	// 建立阻塞状态
	select {}
}

// NewServer 初始化 Server 模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		// 使用 utils.GlobalObject 替换
		ServerName: utils.GlobalObject.Name,
		IpVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		Router:     nil,
	}
	return s
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Router Add Success!!")
}
