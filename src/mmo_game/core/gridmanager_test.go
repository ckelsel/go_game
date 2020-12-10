package core

import (
	"fmt"
	"testing"
)

func TestNewGridManager(t *testing.T) {
	mgr := NewGridManager(0, 250, 5, 0, 250, 5)
	fmt.Println(mgr.String())
}

func TestGetSurroundGrids(t *testing.T) {
	mgr := NewGridManager(0, 250, 5, 0, 250, 5)

	for gid := range mgr.Grids {
		grids := mgr.GetSurroundGrids(gid)
		fmt.Println("gid: ", gid, "grids len: ", len(grids))

		gids := make([]int, 0, len(grids))
		for _, grid := range grids {
			gids = append(gids, grid.ID)
		}

		fmt.Println(" surrounding grids ", gids)
	}
}
