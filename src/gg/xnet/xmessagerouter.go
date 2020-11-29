package xnet

import (
	"fmt"
	"gg/utils"
	"gg/xiface"
)

// XMessageRouter 消息处理抽象层
type XMessageRouter struct {
	// msgid 对应的回调
	Routers map[uint32]xiface.IXRouter

	// 消息队列
	TaskQueue []chan xiface.IXRequest

	// WorkerPool
	WorkerPoolSize uint32
}

// NewXMessageRouter 初始化
func NewXMessageRouter() (router *XMessageRouter) {
	return &XMessageRouter{
		Routers:        make(map[uint32]xiface.IXRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan xiface.IXRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

// Handle 调度Router处理消息
func (mh *XMessageRouter) Handle(request xiface.IXRequest) {
	router, error := mh.Routers[request.GetMsgID()]
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
	_, error := mh.Routers[id]
	if error {
		fmt.Println("ignore msg id ", id)
		return
	}

	mh.Routers[id] = router
	fmt.Println("register msg id ", id)
}

// StartWorkerPool 启动线程池
func (mh *XMessageRouter) StartWorkerPool() {
	var i uint32
	for i = 0; i < mh.WorkerPoolSize; i++ {
		// 创建消息队列
		mh.TaskQueue[i] = make(chan xiface.IXRequest, utils.GlobalObject.MaxTaskQueueSize)

		// 启动worker
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// StartOneWorker 启动一个线程池
func (mh *XMessageRouter) StartOneWorker(id uint32, queue chan xiface.IXRequest) {
	fmt.Println("Start Worker id ", id)

	for {
		select {
		case request := <-queue:
			mh.Handle(request)
		}
	}
}

// PushMessage 消息添加到消息队列
func (mh *XMessageRouter) PushMessage(request xiface.IXRequest) {
	// ConnectionID连接对应的WorkerID
	id := request.GetConn().GetConnID() % utils.GlobalObject.WorkerPoolSize

	// 添加到对应的消息队列
	mh.TaskQueue[id] <- request

	fmt.Println("Worker ", id, "process msg ", request.GetMsgID())
}
