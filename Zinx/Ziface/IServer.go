package Ziface

// Server 定义一个服务器接口
type Server interface {
	// Start 启动服务器
	Start()

	// Stop 运行服务器
	Stop()

	// Serve 运行服务
	Serve()
}
