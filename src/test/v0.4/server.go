package main

import (
	"fmt"
	"xkm/xiface"
	"xkm/xnet"
)


type EchoRouter struct {
    xnet.BaseRouter
}

func (this *EchoRouter) PreHandle(request xiface.IXRequest) {
    fmt.Println("PreHandle")
}

func (this *EchoRouter) Handle(request xiface.IXRequest) {
    fmt.Println("Handle")
    _, err := request.GetConn().GetTCPConnection().Write(request.GetData())
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
