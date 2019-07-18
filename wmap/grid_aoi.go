package wmap

import (
	"fmt"
	"github.com/libaichuan0707/gamelib/util"
	"log"
)

const (
	GRID_AOI_WIDTH_SIZE  = 50
	GRID_AOI_HEIGHT_SIZE = 50
)

type GridAoi struct {
	util.Rectangle
	EnterObjs map[util.ObjectID]bool
}

func NewGridAoi(rectangle util.Rectangle) *GridAoi{
	return &GridAoi {
		Rectangle: rectangle,
		EnterObjs: make(map[util.ObjectID]bool),
	}
}

func (g *GridAoi) GetID() string {
	return fmt.Sprintf("%v_%v", g.X, g.Y)
}

func (g *GridAoi) AddMapObject(id util.ObjectID) {
	g.EnterObjs[id] = true
	log.Printf("GridAoi AddMapObject GridAoiid:%v, objectid:%v\n",g.GetID(), id)
}

func (g *GridAoi) RemoveMapObject(id util.ObjectID) {
	delete(g.EnterObjs, id)
	log.Printf("GridAoi RemoveMapObject GridAoiid:%v, objectid:%v\n",g.GetID(), id)
}

func (g *GridAoi) GetAllObj() map[util.ObjectID]bool {
	return g.EnterObjs
}

func (g *GridAoi) PrintRange() {
	fmt.Printf("x:%v, y:%v, width:%v, height:%v\n", g.X, g.Y, g.Width, g.Height)
}

