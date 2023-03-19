package main

import (
	"ZinxDemo01/Zinx/ziface"
	"ZinxDemo01/Zinx/znet"
	"fmt"
)

/**
基于 Zinx开发的服务端应用
*/

func main() {
	s := znet.NewServer("[Zinx]")

	// 当前 Zinx 框架添加 Router
	s.AddRouter(1, &PingRouter{})

	// 添加Hello Router
	s.AddRouter(2, &HelloRouter{})

	s.Serve()

}

// PingRouter 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// PreHandle 测试路由
func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	// 读取客户端数据,然后回写
	err := request.GetConnection().SendMsg(1, []byte("....Before Ping "))
	if err != nil {
		fmt.Println("Router PreHandle Write Error: ", err)
	}

}

// Handle 测试路由
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	// 读取客户端数据,然后回写

	fmt.Println("Recv Form Client:  MsgID: ", request.GetMsgID(), ",MsgData: ", string(request.GetData()))

	err := request.GetConnection().SendMsg(2, []byte("....Ping....Ping....Ping...."))
	if err != nil {
		fmt.Println("Router Handle SendMsg Error: ", err)
	}
}

// PostHandle 测试路由
func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	err := request.GetConnection().SendMsg(3, []byte(" After Ping...."))
	if err != nil {
		fmt.Println("Router PostHandle Write Error: ", err)
	}
}

// HelloRouter 自定义路由
type HelloRouter struct {
	znet.BaseRouter
}

// Handle 测试路由 返回 Hello
func (hr *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router HelloHandle...")
	sprintf := fmt.Sprintf("Hello, %s", request.GetConnection().GetTCPConnection().RemoteAddr().String())
	fmt.Println("SprintF: ", sprintf)
	err := request.GetConnection().SendMsg(4, []byte(sprintf))
	if err != nil {
		fmt.Println("Router HelloHandle Write Error: ", err)
	}
}
