package znet

import (
	"ZinxDemo01/Zinx/ziface"
	"fmt"
	"strconv"
)

/*
消息处理模块
*/

type MsgHandle struct {
	// 存放每一个MsgID对应的处理方法
	Apis map[uint32]ziface.IRouter

	// 负责处理Worker取任务的消息队列

	// 业务工作数量 worker 数量

}

// NewMsgHandle 创建消息处理对象
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// DoMsgHandler 调度  执行对应的Router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	// 从 Request中找到 MsgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("Api MsgID: ", request.GetMsgID(), " Not Found")
		return
	}
	// 根据MsgID执行对应Router业务
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	// 判断当前msg绑定的Router是否存在
	if _, ok := mh.Apis[msgID]; ok {
		// ID 已经注册了
		panic("Repeat API, MsgID: " + strconv.Itoa(int(msgID)))
	}

	// 添加Msg和API的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add Api MsgID SUCCESS: ", msgID)
}
