package ziface

/*
数据封装接口 解决TCP粘包问题
使用TLV协议来实现封包解包解析
*/

type IDataPack interface {
	// GetHeadLen 获取包长度的方法
	GetHeadLen() uint32
	// Pack 封包方法
	Pack(msg IMessage) ([]byte, error)
	// Unpack 解包方法
	Unpack([]byte) (IMessage, error)
}
