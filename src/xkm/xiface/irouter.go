package xiface

type IRouter interface {
	// 处理业务之前的回调
	PreHandle(request IRequest)

	// 处理业务的主回调
	Handle(request IRequest)

	// 处理业务之后的回调
	PostHandle(request IRequest)
}