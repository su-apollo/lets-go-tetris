package game

type Color uint32
type Cell bool

const (
	o Cell = true
	x Cell = false
)

type Block interface {
	GetCells() [][]Cell
	GetPosition() (int, int)
	GetColor() Color
}

type Board interface {
	GetCells() [][]Cell
	GetColors() [][]Color
}
