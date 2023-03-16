package znet

/*
将请求的消息封装到一个 Message 中
*/

type Message struct {
	ID      uint32 // 消息ID
	DataLen uint32 //消息长度
	Data    []byte // 消息内容
}

// NewMessage 创建Message的方法
func NewMessage(ID uint32, data []byte) *Message {
	return &Message{
		ID:      ID,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// GetMsgID 获取消息ID
func (m *Message) GetMsgID() uint32 {
	return m.ID
}

// GetMsgLen 获取消息长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

// GetData 获取消息内容
func (m *Message) GetData() []byte {
	return m.Data
}

// SetMsgID 设置消息ID
func (m *Message) SetMsgID(id uint32) {
	m.ID = id
}

// SetData 设置消息内容
func (m *Message) SetData(data []byte) {
	m.Data = data
}

// SetDataLen 设置消息长度
func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}
