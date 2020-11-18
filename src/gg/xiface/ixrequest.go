package xiface

// IXRequest 接口
type IXRequest interface {
	// 得到当前链接
	GetConn() IXConnection

	// 得到请求的数据
	GetMsgData() []byte

	// 得到请求的数据ID
	GetMsgID() uint32
}
