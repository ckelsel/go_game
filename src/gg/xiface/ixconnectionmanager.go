package xiface

// IXConnectionManager 连接管理抽象模块
type IXConnectionManager interface {
	// 添加
	Add(conn IXConnection)

	// 删除
	Remove(conn IXConnection)

	// 根据connID获取连接
	Get(connID uint32) (IXConnection, error)

	// 连接总数
	Length() int

	// 清除所有连接
	Clear()
}
