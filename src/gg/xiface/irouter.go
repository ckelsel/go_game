package xiface

type IRouter interface {
	// 处理业务之前的回调
	PreHandle(request IXRequest)

	// 处理业务的主回调
	Handle(request IXRequest)

	// 处理业务之后的回调
	PostHandle(request IXRequest)
}
