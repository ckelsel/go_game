package main

import (
	"fmt"
	"xkm/xiface"
	"xkm/xnet"
)


type EchoRouter struct {
    xnet.BaseRouter
}

func (this *EchoRouter) PreHandle(request xiface.IRequest) {
    fmt.Println("PreHandle")
}

func (this *EchoRouter) Handle(request xiface.IRequest) {
    fmt.Println("Handle")
    _, err := request.GetConn().GetTCPConnection().Write(request.GetData())
    if err != nil {
        fmt.Println("Write err ", err)
    }
}

func (this *EchoRouter) PostHandle(request xiface.IRequest) {
    fmt.Println("PostHandle")
}

func main() {

    s := xnet.NewServer("v0.3")

    router := EchoRouter{}

    s.AddRouter(&router)

    s.Run()
}
