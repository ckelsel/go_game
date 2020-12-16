package main

import (
	"gg/xiface"
	"gg/xnet"
	"mmo_game/core"
)

func OnConnect(conn xiface.IXConnection) {

	player := core.NewPlayer(conn)

	player.SyncPid()

	player.BroadCastStartPosition()

	core.WorldManagerObj.AddPlayer(player)
}

func OnDisconnect(conn xiface.IXConnection) {
	// core.WorldManagerObj.RemovePlayerByPid(player.Pid)
}

func main() {
	s := xnet.NewXServer()

	s.AddOnConnectionStartCallBack(OnConnect)
	s.AddOnConnectionStopCallBack(OnDisconnect)

	s.Start()
}
