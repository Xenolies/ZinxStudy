package znet

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

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

}
