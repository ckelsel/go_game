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
