package main

import (
	"fmt"
	"gg/xiface"
	"gg/xnet"
)

type EchoRouter struct {
	xnet.XBaseRouter
}

func (this *EchoRouter) PreHandle(request xiface.IXRequest) {
	fmt.Println("PreHandle MsgID ", request.GetMsgID(), "MsgData ", request.GetMsgData())
}

func (this *EchoRouter) Handle(request xiface.IXRequest) {
	fmt.Println("Handle")

	err := request.GetConn().SendMsg(request.GetMsgID(), request.GetMsgData())
	if err != nil {
		fmt.Println("Write err ", err)
	}
}

func (this *EchoRouter) PostHandle(request xiface.IXRequest) {
	fmt.Println("PostHandle")
}

func main() {

	s := xnet.NewXServer()

	router := EchoRouter{}

	s.AddRouter(&router)

	s.Run()
}
