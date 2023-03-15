package ziface

/*
将请求的消息封装到一个 Message 中
抽象接口
*/

type IMessage interface {
	// GetMsgID 获取消息ID
	GetMsgID() uint32
	// GetMsgLen 获取消息长度
	GetMsgLen() uint32
	// GetData 获取消息内容
	GetData() []byte

	// SetMsgID 设置消息ID
	SetMsgID(uint32)
	// SetData 设置消息内容
	SetData([]byte)
	// SetDataLen 设置消息长度
	SetDataLen(uint32)
}
