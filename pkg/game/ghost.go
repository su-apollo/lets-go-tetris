package game

var ghostMask = &Color{0xff, 0xff, 0xff, 0x77}

type ghost struct {
	cells [][]Cell
	x, y  int
	color Color
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

func (g *ghost) init(m *matrix, t *tetromino) {
	g.x = t.x
	g.y = t.y
	g.cells = t.GetCells()
	color := t.GetColor()
	g.color = Color{color.R & ghostMask.R, color.G & ghostMask.G, color.B & ghostMask.B, color.A & ghostMask.A}

	drop := true
	for drop {
		g.y++
		drop = !m.collide(g)
	}
	g.y--
}
