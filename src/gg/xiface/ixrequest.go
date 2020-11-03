package xiface

// IXRequest 接口
type IXRequest interface {
	// 得到当前链接
	GetConn() IConnection

	// 得到请求的数据
	GetData() []byte
}
