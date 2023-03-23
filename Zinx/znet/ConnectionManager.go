package znet

import (
	"ZinxStudy/Zinx/ziface"
	"errors"
	"fmt"
	"sync"
)

/*
链接管理模块
*/

type ConnManager struct {
	connections map[uint32]ziface.IConnection //管理链接的集合

	connLock sync.RWMutex // 读写互斥锁

}

// NewConnManager 创建当前链接方法
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
		//connLock: sync.RWMutex{}
	}

}

// Add 添加链接
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	// 保护共享资源Map, 添加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	// 将conn加入到ConnManager中
	connMgr.connections[conn.GetConnID()] = conn
	fmt.Println("ConnID: ", conn.GetConnID(), " ADD TO ConnManager!! Now Conn: ", connMgr.Len())

}

// Remove 删除链接
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	// 保护共享资源Map, 添加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	// 删除链接信息
	delete(connMgr.connections, conn.GetConnID())
	fmt.Println("ConnID: ", conn.GetConnID(), " REMOVE FORM ConnManager!! Now Conn: ", connMgr.Len())

}

// Get 根据connID 获取链接
func (connMgr *ConnManager) Get(ConnID uint32) (ziface.IConnection, error) {
	// 保护共享资源Map, 添加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[ConnID]; ok {
		return conn, nil
	} else {

		return nil, errors.New("conn NOT FOUND FORM ConnManager!! ")
	}
}

// Len 获取当前连接总数
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

// ClearConn 关闭全部链接
func (connMgr *ConnManager) ClearConn() {
	// 保护共享资源Map, 添加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//删除Conn并且停止Conn工作
	for ConnID, conn := range connMgr.connections {
		// 停止conn
		conn.Stop()
		// 删除conn
		delete(connMgr.connections, ConnID)
	}

	fmt.Println("CLEAR All Connections !! Now Conn: ", connMgr.Len())

}
