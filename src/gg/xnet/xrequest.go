package xnet

import (
	"gg/xiface"
)

// XRequest 链接封装为一个请求
type XRequest struct {
	// 链接
	Conn xiface.IConnection

	// 客户端发送过来的数据
	Data []byte
}

// GetConn 得到当前链接
func (r *XRequest) GetConn() xiface.IConnection {
	return r.Conn
}

// GetData 得到请求的数据
func (r *XRequest) GetData() []byte {
	return r.Data
}
