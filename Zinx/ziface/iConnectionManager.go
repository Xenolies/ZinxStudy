package ziface

/*
链接管理器接口
 防止内存泄露
*/

type IConnectionManager interface {
	// Add 添加链接
	Add(conn IConnection)

	// Remove 删除链接
	Remove(conn IConnection)

	// Get 根据connID 获取链接
	Get(ConnID uint32) (IConnection, error)

	// Len 获取当前连接总数
	Len() int

	// ClearConn 关闭全部链接
	ClearConn()
}
