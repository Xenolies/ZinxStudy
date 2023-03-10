package znet

import (
	"ZinxDemo01/Zinx/Ziface"
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
}

func (s *Server) Start() {
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

		fmt.Println("Start Zinx Server Success! ", s.ServerName, "is Listening")

		// 阻塞等待客户端链接和处理客户端链接业务(读写)
		for {
			tcpConn, err := tcpListener.AcceptTCP()
			if err != nil {
				fmt.Println("tcpListener.AcceptTCP Error : ", err)
				continue
			}

			// 客户端链接后的读写操作
			// 做一个最基本的512字节长度的回显业务

			go func() {
				for {
					buf := make([]byte, 512)
					read, err := tcpConn.Read(buf)
					if err != nil {
						fmt.Println("tcpConn.Read Error : ", err)
						continue
					}

					fmt.Printf("Zinx Server Read: %s\n", buf[:read])

					// 回显

					if _, err := tcpConn.Write(buf[:read]); err != nil {
						fmt.Println("tcpConn.Write Error: ", err)
						continue
					}

				}

			}()

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
func NewServer(name string) Ziface.Server {
	s := &Server{
		ServerName: name,
		IpVersion:  "tcp4",
		IP:         "0.0.0.0",
		Port:       8899,
	}
	return s
}
