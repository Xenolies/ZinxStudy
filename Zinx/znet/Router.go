package znet

import (
	"ZinxStudy/Zinx/ziface"
)

// BaseRouter 实现 Router时,线嵌入 BaseRouter基类.然后根据需要对这个基类进行重写
// 实现接口隔离
type BaseRouter struct {
}

// PreHandle 处理Conn业务之前的钩子方法 Hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

// Handle 处理 Conn业务的主方法 Hook
func (br *BaseRouter) Handle(request ziface.IRequest) {
}

// PostHandle 处理Conn 业务之后的钩子方法 Hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {
}
