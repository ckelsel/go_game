package main

import "xkm/xnet"


func main() {

    s := xnet.NewServer("v0.1")

    s.Run()
}
