package wmap

import (
	"fmt"
	"github.com/libaichuan0707/gamelib/util"
)

type GameMap struct {
	AoiMgr      *GridAoiMgr
	GameObjects map[util.ObjectID]IMapObject
}

func NewGameMap(width, height float32) *GameMap {
	return &GameMap {
		AoiMgr: NewGridAoiMgr(width, height),
		GameObjects: make(map[util.ObjectID]IMapObject, 0),
	}
}

func (g *GameMap) OnObjAoiStatuChange(enterID util.ObjectID, observerIDs map[util.ObjectID]bool, status int) {
	for observerID, _ := range observerIDs {
		if status == util.AOI_STATUS_ENTER {
			g.OnObjEnter(enterID, observerID)
		}else if status == util.AOI_STATUS_LEAVE {
			g.OnObjLeave(enterID, observerID)
		}
	}
}

func (g *GameMap) OnObjEnter(enterID, observerID util.ObjectID) {
	if enterID == observerID {
		fmt.Printf("OnObjEnter same enterObjID:%v\n", enterID)
		return
	}
	fmt.Printf("OnObjEnter enterObjID:%v, notifyObjID:%v\n", enterID, observerID)
}

func (g *GameMap) OnObjLeave(enterID, observerID util.ObjectID) {
	if enterID == observerID {
		fmt.Printf("OnObjLeave same enterObjID:%v\n", enterID)
		return
	}
	fmt.Printf("OnObjEnter OnObjLeave:%v, notifyObjID:%v\n", enterID, observerID)
}

func (g *GameMap) AddMapObject(o IMapObject) {
	newGridID, ok := g.AoiMgr.AddMapObject(o.GetPoint().X, o.GetPoint().Y, o.GetMapObjectID())
	if !ok {
		return
	}

	newNears := g.AoiMgr.GetNearGridIDs(newGridID)
	for k, _  := range newNears {
		if grid, ok := g.AoiMgr.GetGrid(k); ok {
			g.OnObjAoiStatuChange(o.GetMapObjectID(), grid.GetAllObj(), util.AOI_STATUS_ENTER)
		}
	}
}

func (g *GameMap) RemoveMapObject(o IMapObject) {
	newGridID, ok := g.AoiMgr.RemoveMapObject(o.GetPoint().X, o.GetPoint().Y, o.GetMapObjectID())
	if !ok {
		return
	}

	newNears := g.AoiMgr.GetNearGridIDs(newGridID)
	for k, _  := range newNears {
		if grid, ok := g.AoiMgr.GetGrid(k); ok {
			g.OnObjAoiStatuChange(o.GetMapObjectID(), grid.GetAllObj(), util.AOI_STATUS_LEAVE)
		}
	}
}