package ziface

/*
 消息管理模块接口
*/

type IMsgHandler interface {
	// DoMsgHandler 调度  执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)

	// AddRouter 为消息添加具体的处理逻辑
	AddRouter(msgID uint32, router IRouter)

	// 启动Worker工作池
	StartWorkPool()
}
