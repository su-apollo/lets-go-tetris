package game

type Block interface {
	GetCells() []cell
	GetPosition() (int, int)
	GetColor() uint32
}

type Board interface {
	GetCells() []cell
	GetColors() []uint32
}
