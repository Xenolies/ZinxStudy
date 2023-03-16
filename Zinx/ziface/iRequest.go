package ziface

// IRequest 接口 将客户端请求的练级信息和请求数据封装到一个Request中
type IRequest interface {
	// GetConnection 得到当前链接
	GetConnection() IConnection
	// GetData 得到请求的消息
	GetData() []byte
	// GetMsgID 获取消息ID
	GetMsgID() uint32
}
