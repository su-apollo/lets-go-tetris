package game

type ghost struct {
	cells [][]Cell
	x, y  int
	shape Shape
}

func (g ghost) Cells() [][]Cell {
	return g.cells
}

func (g ghost) Position() (int, int) {
	return g.x, g.y
}

func (g ghost) Shape() Shape {
	return g.shape
}

func (g *ghost) init(board Board, block Block) {
	g.x, g.y = block.Position()
	g.cells = block.Cells()
	g.shape = block.Shape()

	drop := true
	for drop {
		g.y++
		drop = !board.Collide(g)
	}
	g.y--
}
