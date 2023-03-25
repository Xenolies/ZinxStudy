package main

import (
	"ZinxStudy/Zinx/ziface"
	"ZinxStudy/Zinx/znet"
	"fmt"
)

/**
基于 Zinx开发的服务端应用
*/

// DoConnectionBegin 创建链接之后的Hook函数
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("---> DoConnBegin Hook is Called...")
	if err := conn.SendMsg(202, []byte("DoConnBegin Hook is CALLED!!")); err != nil {
		fmt.Println(err)
	}
}

// DoConnectionLost 创建链接之后的Hook函数
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("---> DoConnStop Hook is Called...")
	fmt.Println("connID: ", conn.GetConnID(), "is Lost....")

}

func main() {
	s := znet.NewServer("[Zinx]")

	// 设置用户创建链接后之调用的 Hook 函数
	s.SetOnConnStart(DoConnectionBegin)
	// 设置用户销毁链接前调用的 Hook 函数
	s.SetOnConnStop(DoConnectionLost)

	// 当前 Zinx 框架添加 Router
	s.AddRouter(1, &PingRouter{})
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
	err := request.GetConnection().SendMsg(1, []byte("....Before Ping...."))
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
	err := request.GetConnection().SendMsg(3, []byte("....After Ping...."))
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
