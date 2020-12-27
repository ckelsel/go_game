package main

import (
	"gg/xiface"
	"gg/xnet"
	"mmo_game/apis"
	"mmo_game/core"
)

func OnConnect(conn xiface.IXConnection) {

	player := core.NewPlayer(conn)

	player.SyncPid()

	player.BroadCastStartPosition()

	conn.SetProperty("pid", player.Pid)

	core.WorldManagerObj.AddPlayer(player)
}

func OnDisconnect(conn xiface.IXConnection) {
	pid, _ := conn.GetProperty("pid")

	player := core.WorldManagerObj.GetPlayerByPid(pid.(int32))

	player.Offline()

}

func main() {
	s := xnet.NewXServer()

	s.AddOnConnectionStartCallBack(OnConnect)
	s.AddOnConnectionStopCallBack(OnDisconnect)

	s.AddRouter(2, &apis.WorldChatAPI{})

	s.Start()
}
