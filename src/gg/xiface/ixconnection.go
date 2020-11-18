package xiface

import (
	"net"
)

// IXConnection 抽象的客户端链接
type IXConnection interface {
	// Start 启动链接，开始工作
	Start()

	// Stop 停止链接
	Stop()

	// GetTCPConnection 获取当前链接绑定的socket
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前链接的ID
	GetConnID() uint32

	// RemoteAddr 获取远程客户端的TCP状态，IP Port
	RemoteAddr() net.Addr

	// Send 发送数据
	Send(data []byte) error

	// SendMsg 发送TLV消息
	SendMsg(id uint32, data []byte) error
}

// HandleFunc 定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
