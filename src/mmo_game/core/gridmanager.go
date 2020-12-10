package core

import "fmt"

// GridManager AOI区域管理模块
type GridManager struct {
	// 区域的左边界坐标
	MinX int
	// 区域的右边界坐标
	MaxX int
	// X方向格子的个数
	X int
	// 区域的上边界坐标
	MinY int
	// 区域的下边界坐标
	MaxY int
	// Y方向格子的个数
	Y int
	// 格子集合
	Grids map[int]*Grid
}

func NewGridManager(minx int, maxx int, x int, miny int, maxy int, y int) *GridManager {
	mgr := &GridManager{
		MinX:  minx,
		MaxX:  maxx,
		X:     x,
		MaxY:  maxy,
		MinY:  miny,
		Y:     y,
		Grids: make(map[int]*Grid),
	}

	for yy := 0; yy < y; yy++ {
		for xx := 0; xx < x; xx++ {
			// 格子编号
			id := yy*x + xx

			mgr.Grids[id] = NewGrid(id,
				mgr.MinX+xx*mgr.GridWidth(),
				mgr.MinX+(xx+1)*mgr.GridWidth(),
				mgr.MinY+yy*mgr.GridHeight(),
				mgr.MinY+(yy+1)*mgr.GridHeight(),
			)
		}
	}
	return mgr
}

// GridWidth 格子的X宽度
func (m *GridManager) GridWidth() int {
	return (m.MaxX - m.MinX) / m.X
}

// GridHeight 格子的Y宽度
func (m *GridManager) GridHeight() int {
	return (m.MaxY - m.MinY) / m.Y
}

func (m *GridManager) String() string {
	s := fmt.Sprintf("GridManager:\nMinX %d, MaxX %d, MinY %d, MaxY %d, X %d, Y %d",
		m.MinX, m.MaxX, m.MinY, m.MaxY, m.X, m.Y)
	for _, grid := range m.Grids {
		s += fmt.Sprintln(grid)
	}

	return s
}

// GetSurroundGrids 根据GID获取周围的所有格子
func (m *GridManager) GetSurroundGrids(gid int) (grids []*Grid) {
	if _, success := m.Grids[gid]; !success {
		return
	}

	grids = append(grids, m.Grids[gid])

	idx := gid % m.X

	if idx > 0 {
		grids = append(grids, m.Grids[gid-1])
	}

	if idx < m.X-1 {
		grids = append(grids, m.Grids[gid+1])
	}

	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.ID)
	}

	for _, v := range gidsX {
		idy := v / m.Y

		if idy > 0 {
			grids = append(grids, m.Grids[v-m.X])
		}

		if idy < m.Y-1 {
			grids = append(grids, m.Grids[v+m.X])
		}
	}

	return grids
}

// GetGidByPosition 通过x,y坐标得到当前的GID格子编号
func (m *GridManager) GetGidByPosition(x, y float32) int {
	idx := (int(x) - m.MinX) / m.GridWidth()
	idy := (int(y) - m.MinY) / m.GridHeight()

	return idy*m.X + idx
}

// GetAllPlayerIDByPosition 通过x, y坐标得到周边九宫格内全部的playerids
func (m *GridManager) GetAllPlayerIDByPosition(x, y float32) (players []int) {
	gid := m.GetGidByPosition(x, y)

	grids := m.GetSurroundGrids(gid)

	for _, v := range grids {
		v.GetPlayerIDs(players)
	}

	return players
}

// AddPlayerIDToGrid 添加一个playerid到格子
func (m *GridManager) AddPlayerIDToGrid(pid, gid int) {
	m.Grids[gid].Add(pid)
}

// RemovePlayerIDFromGrid 删除一个格子内的playerid
func (m *GridManager) RemovePlayerIDFromGrid(pid, gid int) {
	m.Grids[gid].Remove(pid)
}

// GetAllPlayerIDByGid 通过gid获取全部的playerid
func (m *GridManager) GetAllPlayerIDByGid(gid int) (players []int) {
	m.Grids[gid].GetPlayerIDs(players)

	return players
}

// AddPositionToGrid 通过坐标讲player添加到格子
func (m *GridManager) AddPositionToGrid(x, y float32, gid int) {
	m.Grids[gid].Add(m.GetGidByPosition(x, y))
}

// RemovePositionFromGrid 通过坐标从格子删除player
func (m *GridManager) RemovePositionFromGrid(x, y float32, gid int) {
	m.Grids[gid].Remove(m.GetGidByPosition(x, y))
}
