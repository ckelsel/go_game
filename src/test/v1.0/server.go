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

func OnStart(conn xiface.IXConnection) {
	fmt.Println("OnStart")

    conn.SetProperty("date", "2020-12-07")
}

func OnStop(conn xiface.IXConnection) {
	fmt.Println("OnStop")

    if date, err:= conn.GetProperty("date"); err == nil {
        fmt.Println("date ", date)
    }
}

func main() {

	s := xnet.NewXServer()

	router := EchoRouter{}

	s.AddRouter(0, &router)
	s.AddOnConnectionStartCallBack(OnStart)
	s.AddOnConnectionStopCallBack(OnStop)
	s.Run()
}
