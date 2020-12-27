package core

import (
	"fmt"
	"gg/xiface"
	"math/rand"
	"mmo_game/pb"
	"sync"

	"google.golang.org/protobuf/proto"
)

// Player 玩家
type Player struct {
	Pid  int32
	Conn xiface.IXConnection

	X float32 //x坐标
	Y float32 // 高度
	Z float32 // y坐标
	V float32 // 旋转角度 0-360
}

var PidGen int32 = 1
var PidGenMutex sync.Mutex

func NewPlayer(conn xiface.IXConnection) *Player {
	PidGenMutex.Lock()
	id := PidGen
	PidGen++
	PidGenMutex.Unlock()

	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(150 + rand.Intn(20)),
		V:    0,
	}

	return p
}

func (p *Player) SendMessage(msgid uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal err: ", err)
		return
	}

	if p.Conn == nil {
		fmt.Println("connection lost")
		return
	}

	if err := p.Conn.SendMsg(msgid, msg); err != nil {
		fmt.Println("SendMsg failed")
		return
	}
}

func (p *Player) SyncPid() {
	data := &pb.SyncPid{
		Pid: p.Pid,
	}
	p.SendMessage(1, data)
}

func (p *Player) BroadCastStartPosition() {
	data := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMessage(200, data)
}

func (p *Player) Talk(content string) {
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}

	players := WorldManagerObj.GetAllPlayer()

	for _, player := range players {
		player.SendMessage(200, msg)
	}
}

//获得当前玩家的AOI周边玩家信息
func (p *Player) GetSurroundingPlayers() []*Player {
	//得到当前AOI区域的所有pid
	pids := WorldManagerObj.GM.GetAllPlayerIDByPosition(p.X, p.Z)

	//将所有pid对应的Player放到Player切片中
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldManagerObj.GetPlayerByPid(int32(pid)))
	}

	return players
}

func (p *Player) Offline() {
	players := p.GetSurroundingPlayers()

	msg := &pb.SyncPid{
		Pid: p.Pid,
	}

	for _, player := range players {
		player.SendMessage(201, msg)
	}

	WorldManagerObj.GM.RemovePositionFromGrid(p.X, p.Z, (int)(p.Pid))
	WorldManagerObj.RemovePlayerByPid(p.Pid)
}
