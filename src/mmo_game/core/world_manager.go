package core

import "sync"

const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

// WorldManager 游戏管理模块
type WorldManager struct {
	GM *GridManager

	Players map[int32]*Player

	PlayersMutex sync.Mutex
}

var WorldManagerObj *WorldManager

func init() {
	WorldManagerObj = &WorldManager{
		GM:      NewGridManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
		Players: make(map[int32]*Player),
	}
}

// 添加一个玩家
func (wm *WorldManager) AddPlayer(player *Player) {
	wm.PlayersMutex.Lock()
	wm.Players[player.Pid] = player
	wm.PlayersMutex.Unlock()

	wm.GM.AddPositionToGrid(player.X, player.Z, int(player.Pid))
}

// 删除一个玩家
func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	player := wm.Players[pid]

	wm.GM.RemovePositionFromGrid(player.X, player.Z, int(pid))

	wm.PlayersMutex.Lock()
	delete(wm.Players, pid)
	wm.PlayersMutex.Unlock()
}

// 通过玩家ID查询Player对象
func (wm *WorldManager) GetPlayerByPid(pid int32) *Player {
	wm.PlayersMutex.Lock()
	defer wm.PlayersMutex.Unlock()

	return wm.Players[pid]
}

// 获取全部在线的玩家
func (wm *WorldManager) GetAllPlayer(pid int32) []*Player {
	wm.PlayersMutex.Lock()
	defer wm.PlayersMutex.Unlock()

	players := make([]*Player, 0)

	for _, p := range wm.Players {
		players = append(players, p)
	}

	return players
}
