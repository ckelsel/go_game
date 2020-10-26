package xnet

import (
	"xkm/xiface"
)

type Request struct {
	// 链接
	Conn xiface.IConnection 

	// 客户端发送过来的数据
	Data []byte 
}


// 得到当前链接
func (r *Request)GetConn() xiface.IConnection {
	return r.Conn;
}

// 得到请求的数据
func (r *Request)GetData() []byte {
	return r.Data
}