package xiface

// IXServer 抽象的服务器
type IXServer interface {
	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// Run 运行服务器
	Run()

	// AddRouter 用户自定义的路由功能
	AddRouter(msgid uint32, router IXRouter)

	// 设置连接建立后的回调
	AddOnConnectionStartCallBack(onstart func(conn IXConnection))

	// 设置连接断开前的回调
	AddOnConnectionStopCallBack(onstop func(conn IXConnection))

	// 获取连接管理器
	GetConnectionManager() (cm IXConnectionManager)

	// 连接建立后的回调
	OnConnectionStart(conn IXConnection)

	// 连接断开前的回调
	OnConnectionStop(conn IXConnection)
}
