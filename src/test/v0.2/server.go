package main

import "xkm/xnet"


func main() {

    s := xnet.NewXServer("v0.2")

    s.Run()
}
