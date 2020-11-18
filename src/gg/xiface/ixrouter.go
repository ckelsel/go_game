package xiface

// IXRouter 处理客户端消息的路由
type IXRouter interface {
	// PreHandle 处理业务之前的回调
	PreHandle(request IXRequest)

	// Handle 处理业务的主回调
	Handle(request IXRequest)

	// PostHandle 处理业务之后的回调
	PostHandle(request IXRequest)
}
