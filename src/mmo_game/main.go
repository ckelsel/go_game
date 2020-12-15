package main

import (
	"fmt"
	"gg/xiface"
	"gg/xnet"
	"mmo_game/core"
)

func OnStart(conn xiface.IXConnection) {

	fmt.Println("player arrived")
	player := core.NewPlayer(conn)
	fmt.Println("player arrived3")

	player.SyncPid()
	fmt.Println("player arrived4")

	player.BroadCastStartPosition()
	fmt.Println("player arrived2")

}

func main() {
	s := xnet.NewXServer()

	s.AddOnConnectionStartCallBack(OnStart)

	s.Start()
}
