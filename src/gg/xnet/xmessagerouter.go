package xnet

import (
	"fmt"
	"gg/xiface"
)

// XMessageRouter 消息处理抽象层
type XMessageRouter struct {
	routers map[uint32]xiface.IXRouter
}

// NewXMessageRouter 初始化
func NewXMessageRouter() (router *XMessageRouter) {
	return &XMessageRouter{
		routers: make(map[uint32]xiface.IXRouter),
	}
}

// Handle 调度Router处理消息
func (mh *XMessageRouter) Handle(request xiface.IXRequest) {
	router, error := mh.routers[request.GetMsgID()]
	if !error {
		fmt.Println("router not find, msg id ", request.GetMsgID())
		return
	}

	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

// AddRouter 为特定的消息添加Router
func (mh *XMessageRouter) AddRouter(id uint32, router xiface.IXRouter) {
	_, error := mh.routers[id]
	if error {
		fmt.Println("ignore msg id ", id)
		return
	}

	mh.routers[id] = router
	fmt.Println("register msg id ", id)
}
