package xiface

// IXMessageRouter 消息处理抽象层
type IXMessageRouter interface {
	// Handle 调度Router处理消息
	Handle(request IXRequest)

	// AddRouter 为特定的消息添加Router
	AddRouter(id uint32, router IXRouter)

	// PushMessage 消息添加到消息队列
	PushMessage(request IXRequest)

	// StartWorkerPool 启动线程池
	StartWorkerPool()
}
