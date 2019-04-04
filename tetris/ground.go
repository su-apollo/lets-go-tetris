package tetris

type Ground struct {
	X     int
	Y     int
	Cells []Cell
}

func NewGround(x, y int) *Ground {
	return &Ground{x, y, make([]Cell, x*y)}
}

func (ground *Ground) Clear() {
	for i := 0; i < ground.X*ground.Y; i++ {
		ground.Cells[i].Filled = false
	}
}

func (ground *Ground) Draw() {
}
