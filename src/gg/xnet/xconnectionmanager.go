package xnet

import (
	"errors"
	"fmt"
	"gg/xiface"
	"sync"
)

// XConnectionManager 连接管理模块
type XConnectionManager struct {
	Connections map[uint32]xiface.IXConnection

	Mutex sync.RWMutex
}

// NewXConnectionManager 初始化方法
func NewXConnectionManager() *XConnectionManager {
	return &XConnectionManager{
		Connections: make(map[uint32]xiface.IXConnection),
	}
}

// Add 添加
func (cm *XConnectionManager) Add(conn xiface.IXConnection) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()

	cm.Connections[conn.GetConnID()] = conn

	fmt.Println("connection ", conn.GetConnID(), "add to ConnectionManager, Length ", cm.Length())
}

// Remove 删除
func (cm *XConnectionManager) Remove(conn xiface.IXConnection) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()

	delete(cm.Connections, conn.GetConnID())

	fmt.Println("connection ", conn.GetConnID(), "remove from ConnectionManager, Length ", cm.Length())
}

// Get 根据connID获取连接
func (cm *XConnectionManager) Get(connID uint32) (xiface.IXConnection, error) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()

	conn, success := cm.Connections[connID]
	if success {
		return conn, nil
	} else {
		return nil, errors.New("connection not find")
	}
}

// Length 连接总数
func (cm *XConnectionManager) Length() int {
	return len(cm.Connections)
}

// Clear 清除所有连接
func (cm *XConnectionManager) Clear() {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()

	for connID, conn := range cm.Connections {
		conn.Stop()

		delete(cm.Connections, connID)
	}

	fmt.Println("Clear all connections")
}
