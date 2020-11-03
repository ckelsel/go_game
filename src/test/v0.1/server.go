package main

import "gg/xnet"


func main() {

    s := xnet.NewXServer("v0.1")

    s.Run()
}
