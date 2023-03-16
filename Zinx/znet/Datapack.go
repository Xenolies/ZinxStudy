package znet

import (
	"ZinxDemo01/Zinx/utils"
	"ZinxDemo01/Zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

// DataPack 拆包和封包模块
type DataPack struct {
	Message ziface.IMessage
}

// NewDataPack 拆包封包实例的初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 获取包长度的方法
func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen uint32 + ID uint32 = 8 字节
	return 8
}

// Pack 封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放[]byte的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 将dataLen写进dataBuff
	// 高端对高地址
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}
	// 将MsgID写进dataBuff
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID())
	if err != nil {
		return nil, err
	}
	// 将data写进dataBuff
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// Unpack 解包方法
// 拆包先读头 (长度,类型ID) 再读尾 (内容)
// 先把Head读出来,然后读取Head里面data长度
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建一个从输入二进制读数据的Buff
	dataBuff := bytes.NewBuffer(binaryData)
	// 创建 Message对象 接受解析的数据
	msg := &Message{}

	// 只解压头部信息,得到DataLen 和Message ID信息
	// 读DataLen
	err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}

	// 读Message ID
	err = binary.Read(dataBuff, binary.BigEndian, &msg.ID)
	if err != nil {
		return nil, err
	}
	// 用户在设置中设置了最大包大小,需要判断是否超出最大长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("msgData is too Large")
	}

	// 解析内容

	return msg, nil
}
