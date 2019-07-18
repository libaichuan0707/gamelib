package wmap

import "testing"

func Test_Mapt1(t *testing.T) {
	m := NewGameMap(220, 200)
	all := make([]IMoveMapObject, 0)
	for i := 1; i <= 3; i++ {
		o := NewMoveMapObject(float32(i *40), 10)
		m.AddMapObject(o)
		all = append(all, o)
	}

	for i := 1; i <= 2; i++ {
		o := NewPlayer(float32(i *30), 10, i)
		m.AddMapObject(o)
		all = append(all, o)
	}

	all[4].UpdatePos(m, 140, 10)
}
