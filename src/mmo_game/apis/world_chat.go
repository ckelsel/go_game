package apis

import (
	"fmt"
	"gg/xiface"
	"gg/xnet"
	"mmo_game/core"
	"mmo_game/pb"

	"google.golang.org/protobuf/proto"
)

// WorldChatAPI 世界聊天路由
type WorldChatAPI struct {
	xnet.XBaseRouter
}

func (wc *WorldChatAPI) Handle(request xiface.IXRequest) {
	msg := &pb.Talk{}

	err := proto.Unmarshal(request.GetMsgData(), msg)
	if err != nil {
		fmt.Println("Talk Unmarshal failed, ", err)
		return
	}

	pid, err := request.GetConn().GetProperty("pid")

	player := core.WorldManagerObj.GetPlayerByPid(pid.(int32))

	player.Talk(msg.Content)
}
