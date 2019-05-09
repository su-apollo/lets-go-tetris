package game

const ghostColor = 0x77ffffff

type ghost struct {
	cells []cell
	x, y  int
	color uint32
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

func (g *ghost) GetCells() []cell {
	return g.cells
}

func (g *ghost) GetPosition() (int, int) {
	return g.x, g.y
}

func (g *ghost) GetColor() uint32 {
	return g.color
}
