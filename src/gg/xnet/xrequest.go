package xnet

import (
	"gg/xiface"
)

// XRequest 链接封装为一个请求
type XRequest struct {
	// 链接
	Conn xiface.IXConnection

	// 客户端发送过来的数据
	Message xiface.IXMessage
}

// GetConn 得到当前链接
func (r *XRequest) GetConn() xiface.IXConnection {
	return r.Conn
}

// GetMsgData 得到请求的数据
func (r *XRequest) GetMsgData() []byte {
	return r.Message.GetData()
}

// GetMsgID 得到请求的数据ID
func (r *XRequest) GetMsgID() uint32 {
	return r.Message.GetID()
}
