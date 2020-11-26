package xnet

import (
	"gg/xiface"
)

// XBaseRouter 处理业务的默认路由,默认空实现
// 具体实现，需要重新BaseRouter的方法
type XBaseRouter struct {
}

// PreHandle 处理业务之前的回调
func (b *XBaseRouter) PreHandle(request xiface.IXRequest) {

}

// Handle 处理业务的主回调
func (b *XBaseRouter) Handle(request xiface.IXRequest) {

}

// PostHandle 处理业务之后的回调
func (b *XBaseRouter) PostHandle(request xiface.IXRequest) {

}
