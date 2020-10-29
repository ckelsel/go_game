package xiface

type IXServer interface {
	// 启动服务器
	Start()

	// 停止服务器
	Stop()

	// 运行服务器
	Run()

	// 路由功能
	AddRouter(router IRouter)
}
