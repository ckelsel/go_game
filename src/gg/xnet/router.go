package xnet

import (
	"gg/xiface"
)

// 处理业务的默认路由,默认空实现
// 具体实现，需要重新BaseRouter的方法
type BaseRouter struct {
}

// 处理业务之前的回调
func (b *BaseRouter) PreHandle(request xiface.IXRequest) {

}

// 处理业务的主回调
func (b *BaseRouter) Handle(request xiface.IXRequest) {

}

// 处理业务之后的回调
func (b *BaseRouter) PostHandle(request xiface.IXRequest) {

}
