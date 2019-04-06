package game

type ground struct {
	x, y  int
	cells []cell
}

func (ground *ground) reset() {
	ground.cells = make([]cell, ground.x*ground.y)
}

func (ground *ground) draw() {
}
