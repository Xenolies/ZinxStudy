package Ziface

// IRouter 路由抽象接口  路由里面数据都是IRequest
type IRouter interface {
	// PreHandle 处理Conn业务之前的钩子方法 Hook
	PreHandle(request IRequest)

	// Handle 处理 Conn业务的主方法 Hook
	Handle(request IRequest)

	// PostHandle 处理Conn 业务之后的钩子方法 Hook
	PostHandle(request IRequest)
}
