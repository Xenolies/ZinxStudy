package znet

import "ZinxDemo01/Zinx/ziface"

type Request struct {
	// 已经和客户端建立好的链接
	conn ziface.IConnection

	//客户端请求的数据
	data []byte
}

// GetConnection 获取客户端链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 获取用户端请求的数据
func (r *Request) GetData(data []byte) []byte {
	return r.data
}
