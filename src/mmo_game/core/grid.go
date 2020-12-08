package core

import (
	"fmt"
	"sync"
)

// Grid AOI格子
type Grid struct {
	// 格子ID
	ID int
	// 左边界
	MinX int
	// 右边界
	MaxX int
	// 上边界
	MinY int
	// 下边界
	MaxY int
	// 玩家集合
	PlayerID map[int]bool
	// 玩家集合的互斥锁
	PlayerIDMutex sync.RWMutex
}

// NewGrid 初始化
func NewGrid(id int, minx int, maxx int, miny int, maxy int) *Grid {
	return &Grid{
		ID:       id,
		MinX:     minx,
		MaxX:     maxx,
		MinY:     miny,
		MaxY:     maxy,
		PlayerID: make(map[int]bool),
	}
}

// Add 添加一个玩家
func (g *Grid) Add(playerid int) {
	g.PlayerIDMutex.Lock()
	defer g.PlayerIDMutex.Unlock()

	g.PlayerID[playerid] = true
}

// Remove 删除一个玩家
func (g *Grid) Remove(playerid int) {
	g.PlayerIDMutex.Lock()
	defer g.PlayerIDMutex.Unlock()

	delete(g.PlayerID, playerid)
}

// GetPlayerIDs 获取玩家列表
func (g *Grid) GetPlayerIDs(player []int) {
	g.PlayerIDMutex.Lock()
	defer g.PlayerIDMutex.Unlock()

	for k := range g.PlayerID {
		player = append(player, k)
	}
}

// 打印格子信息
func (g *Grid) String() string {
	return fmt.Sprintf("Grid: id %d, MinX %d, MaxX %d, MinY %d, MaxY %d",
		g.ID, g.MinX, g.MaxX, g.MinY, g.MaxY)
}
