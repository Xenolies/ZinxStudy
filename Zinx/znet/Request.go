package znet

import "ZinxDemo01/Zinx/Ziface"

type Request struct {
	// 已经和客户端建立好的链接
	conn Ziface.IConnection

	//客户端请求的数据
	data []byte
}

// GetConnection 获取客户端链接
func (r *Request) GetConnection() Ziface.IConnection {
	return r.conn
}

// GetData 获取用户端请求的数据
func (r *Request) GetData(data []byte) []byte {
	return r.data
}
