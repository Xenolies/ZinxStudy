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
	AddRouter(msgID uint32, router IRouter)

	// GetConnMgr 获取当前Server的ConnManager
	GetConnMgr() IConnectionManager

	// SetOnConnStart 注册Server创建链接自动调用的 Hook函数
	SetOnConnStart(func(conn IConnection))

	// SetOnConnStop 注册Server销毁链接自动调用的 Hook函数
	SetOnConnStop(func(conn IConnection))

	// CallOnConnStart 调用Server创建链接自动调用的 Hook函数
	CallOnConnStart(connection IConnection)

	// CallOnConnStop 调用Server销毁链接自动调用的 Hook函数
	CallOnConnStop(connection IConnection)
}
