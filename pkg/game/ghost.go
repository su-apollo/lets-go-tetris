package game

const ghostColor = 0x77ffffff

type ghost struct {
	cells [][]Cell
	x, y  int
	color Color
}

func (g *ghost) init(m *matrix, t *tetromino) {
	g.x = t.x
	g.y = t.y
	g.cells = t.GetCells()
	g.color = t.GetColor() & ghostColor

	drop := true
	for drop {
		g.y++
		drop = !m.collide(g)
	}
	g.y--
}

func (g *ghost) GetCells() [][]Cell {
	return g.cells
}

func (g *ghost) GetPosition() (int, int) {
	return g.x, g.y
}

func (g *ghost) GetColor() Color {
	return g.color
}
