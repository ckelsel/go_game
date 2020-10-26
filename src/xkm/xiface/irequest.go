package xiface

type IRequest interface {
	// 得到当前链接
	GetConn() IConnection

	// 得到请求的数据
	GetData() []byte
}