package main

import (
	"ZinxDemo01/Zinx/Ziface"
	"ZinxDemo01/Zinx/znet"
	"fmt"
)

/**
基于 Zinx开发的服务端应用
*/

func main() {
	s := znet.NewServer("[Zinx]")

	// 当前 Zinx 框架添加 Router
	s.AddRouter(&PingRouter{})

	s.Serve()

}

// PingRouter 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// PreHandle 测试路由
func (pr *PingRouter) PreHandle(request Ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Before Ping....  |  "))
	if err != nil {
		fmt.Println("Router PreHandle Write Error: ", err)
	}

}

// Handle 测试路由
func (pr *PingRouter) Handle(request Ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("....Ping....Ping....Ping....  |  "))
	if err != nil {
		fmt.Println("Router Handle Write Error: ", err)
	}
}

// PostHandle 测试路由
func (pr *PingRouter) PostHandle(request Ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After Ping...."))
	if err != nil {
		fmt.Println("Router PostHandle Write Error: ", err)
	}
}
