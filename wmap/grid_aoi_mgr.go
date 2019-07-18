package wmap

import (
	"fmt"
	"github.com/libaichuan0707/gamelib/util"
	"log"
	"math"
)

type GridAoiMgr struct {
	Grids []*GridAoi
	Width, Height float32
	WidthNum, HeightNum int
}

func NewGridAoiMgr(width, height float32) *GridAoiMgr {
	mgr := &GridAoiMgr {
		Grids: make([]*GridAoi, 0),
		Width: width,
		Height: height,
		WidthNum: int(math.Ceil(float64(width) / GRID_AOI_WIDTH_SIZE)),
		HeightNum: int(math.Ceil(float64(height) / GRID_AOI_HEIGHT_SIZE)),
	}

	for h := 0; h < mgr.HeightNum; h++ {
		for w := 0; w < mgr.WidthNum; w++ {
			grid := NewGridAoi(util.Rectangle{
				Point:util.Point{
					X: float32(w*GRID_AOI_WIDTH_SIZE),
					Y: float32(h*GRID_AOI_HEIGHT_SIZE),
				},
				Width: GRID_AOI_WIDTH_SIZE,
				Height: GRID_AOI_HEIGHT_SIZE,
			})

			if w == mgr.WidthNum - 1 {
				grid.Width = width - float32(w) * GRID_AOI_WIDTH_SIZE
			}

			if h == mgr.HeightNum -1 {
				grid.Height = height - float32(h) * GRID_AOI_HEIGHT_SIZE
			}

			mgr.Grids = append(mgr.Grids, grid)
		}
	}

	return mgr
}

func (g *GridAoiMgr) Print() {
	for i, v := range g.Grids {
		fmt.Printf("index:%v  ", i)
		v.PrintRange()
	}
}

func (g *GridAoiMgr) GetGridID(x, y float32) int {
	w := int(x / GRID_AOI_WIDTH_SIZE)
	h := int(y / GRID_AOI_HEIGHT_SIZE)

	return g.GetGridIDByGrid(w, h)
}

func (g *GridAoiMgr) GetGridIDByGrid(widthIndex, heightIndex int) int {
	return heightIndex * g.WidthNum + widthIndex
}

func (g *GridAoiMgr) GetGrid(gridID int) (*GridAoi, bool) {
	if len(g.Grids) <= gridID || gridID < 0{
		return nil, false
	}

	return g.Grids[gridID], true
}

func (g *GridAoiMgr) CheckPos(x, y float32) bool {
	if x >= 0 && x <= g.Width && y >= 0 && y <= g.Height {
		return true
	}

	return false
}

func (g *GridAoiMgr) CheckGridIndex(widthIndex, heighIndex int) bool {
	if widthIndex >= 0 && widthIndex < g.WidthNum && heighIndex >= 0 && heighIndex <= g.HeightNum {
		return true
	}

	return false
}

func (g *GridAoiMgr) CheckGridID(gridID int) bool {
	if gridID < len(g.Grids) && gridID >= 0 {
		return true
	}

	return false
}

func (g *GridAoiMgr) GetNearGridIDs(gridID int) map[int]bool {
	gridIDs := make(map[int]bool)
	if !g.CheckGridID(gridID) {
		return gridIDs
	}

	width := gridID % g.WidthNum
	height := (gridID - width) / g.WidthNum

	log.Printf("GetNearGridIDs gridID:%v, width:%v, height:%v", gridID, width, height)

	for i := -1; i<= 1; i++ {
		for j := -1; j <= 1; j++ {
			if g.CheckGridIndex(width + i, height+j) {
				gridIDs[g.GetGridIDByGrid(width + i, height + j)] = true
			}
		}
	}

	return gridIDs
}

func (g *GridAoiMgr) AddMapObject(x, y float32, id util.ObjectID) (int, bool) {
	if !g.CheckPos(x, y) {
		log.Printf("error pos x:%v, y:%v", x, y)
		return -1, false
	}

	newGridID := g.GetGridID(x, y)

	if !g.CheckGridID(newGridID) {
		log.Printf("error newGridID x:%v, y:%v", x, y)
		return -1, false
	}

	g.Grids[newGridID].AddMapObject(id)

	return newGridID, true
}

func (g *GridAoiMgr) RemoveMapObject(x, y float32, id util.ObjectID) (int, bool)  {
	if !g.CheckPos(x, y) {
		log.Printf("error pos x:%v, y:%v", x, y)
		return -1, false
	}

	newGridID := g.GetGridID(x, y)

	if !g.CheckGridID(newGridID) {
		return -1, false
	}

	g.Grids[newGridID].RemoveMapObject(id)

	return newGridID, true
}

func GetSameMap(oldNears map[int]bool, newNears map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for k, _ := range oldNears {
		if _, ok := newNears[k]; ok {
			result[k] = true
		}
	}

	return result
}