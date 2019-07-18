package wmap

import (
	"github.com/libaichuan0707/gamelib/util"
	"log"
)

type IMapObject interface {
	GetMapObjectID() util.ObjectID
	GetPoint() util.Point
}

type IMoveMapObject interface {
	IMapObject
	UpdatePos(gameMap *GameMap, newX, newY float32)
}

type BaseMapObject struct {
	ID util.ObjectID
	util.Point
}

func NewBaseMapObject(x, y float32) *BaseMapObject {
	m := &BaseMapObject{}
	m.X = x
	m.Y = y
	m.ID = util.GenNewID()
	return m
}

func (m *BaseMapObject) GetMapObjectID() util.ObjectID {
	return m.ID
}

func (m *BaseMapObject) GetPoint() util.Point {
	return m.Point
}

type MoveMapObject struct {
	*BaseMapObject
}

func NewMoveMapObject(x, y float32) *MoveMapObject {
	m := &MoveMapObject{}
	m.BaseMapObject = NewBaseMapObject(x, y)
	return m
}

func (m *MoveMapObject) UpdatePos(gameMap *GameMap, newX, newY float32) {
	g := gameMap.AoiMgr
	if !g.CheckPos(newX, newY) {
		log.Printf("error pos x:%v, y:%v", newX, newY)
		return
	}

	oldGridID := g.GetGridID(m.X, m.Y)
	newGridID := g.GetGridID(newX, newY)

	log.Printf("oldGridID:%v, newGridID:%v\n", oldGridID, newGridID)

	m.X = newX
	m.Y = newY

	if oldGridID != newGridID {
		if g.CheckGridID(oldGridID) {
			g.Grids[oldGridID].RemoveMapObject(m.GetMapObjectID())
		}

		if g.CheckGridID(newGridID) {
			g.Grids[newGridID].AddMapObject(m.GetMapObjectID())
		}

		oldNears := g.GetNearGridIDs(oldGridID)
		newNears := g.GetNearGridIDs(newGridID)

		sameNears := GetSameMap(oldNears, newNears)

		log.Printf("oldNears:%v, newNears:%v, sameNears:%v \n", oldNears, newNears, sameNears)

		for k, _  := range oldNears {
			if _, ok := sameNears[k]; ok {
				continue
			}

			if grid, ok := g.GetGrid(k); ok {
				gameMap.OnObjAoiStatuChange(m.GetMapObjectID(), grid.EnterObjs, util.AOI_STATUS_LEAVE)
			}
		}

		for k, _  := range newNears {
			if _, ok := sameNears[k]; ok {
				continue
			}

			if grid, ok := g.GetGrid(k); ok {
				gameMap.OnObjAoiStatuChange(m.GetMapObjectID(), grid.EnterObjs, util.AOI_STATUS_ENTER)
			}
		}
	}
}

type Player struct {
	*MoveMapObject
	Sex int
}

func NewPlayer(x, y float32, sex int) *Player {
	p := &Player{}
	p.MoveMapObject = NewMoveMapObject(x, y)
	p.Sex = sex
	return p
}