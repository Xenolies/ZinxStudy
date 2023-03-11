package ziface

// IServer Server 定义一个服务器接口
type IServer interface {
	// Start 启动服务器
	Start()

	// Stop 运行服务器
	Stop()

	// Serve 运行服务
	Serve()

	// AddRouter 路由功能 给当前服务注册一个 路由,来处理客户端链接
	AddRouter(router IRouter)
}
