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
