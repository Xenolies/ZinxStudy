package znet

import "ZinxDemo01/Zinx/ziface"

type Request struct {
	// 已经和客户端建立好的链接
	conn ziface.IConnection

	//客户端请求的数据
	msg ziface.IMessage
}

// GetConnection 获取客户端链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 获取用户端请求的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID 获取消息 ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
