package util

type ObjectID int64

const (
	AOI_STATUS_ENTER = 0
	AOI_STATUS_LEAVE = 1
)

type Point struct {
	X, Y float32
}

type Rectangle struct {
	// 左上角坐标
	Point
	Width, Height float32
}