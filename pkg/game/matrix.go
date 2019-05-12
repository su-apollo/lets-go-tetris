package game

var tileColors = []Color{
	0xff353535,
	0xff5D5D5D,
}

type matrix struct {
	width, height int
	cells         [][]Cell
	colors        [][]Color
}

func (m *matrix) reset() {
	m.cells = make([][]Cell, m.height)
	for i := range m.cells {
		m.cells[i] = make([]Cell, m.width)
	}

	m.colors = make([][]Color, m.height)
	for i := range m.colors {
		m.colors[i] = make([]Color, m.width)
	}
}

func (m *matrix) collide(b Block) bool {
	for y, cells := range b.GetCells() {
		for x, cell := range cells {
			if cell {
				cx, cy := b.GetPosition()
				cx += x
				cy += y

				if cx < 0 || m.width <= cx || cy < 0 || m.height <= cy {
					return true
				}

				if m.cells[cy][cx] {
					return true
				}
			}
		}
	}
	return false
}

func (m *matrix) merge(b Block) {
	for y, cells := range b.GetCells() {
		for x, cell := range cells {
			if cell {
				cx, cy := b.GetPosition()
				cx += x
				cy += y

				if 0 <= cx && cx < m.width && 0 <= cy && cy < m.height {
					m.cells[cy][cx] = true
					m.colors[cy][cx] = b.GetColor()
				}
			}
		}
	}
}

func (m *matrix) removeLines() int {
	lines := 0
	for y := 0; y < m.height; y++ {
		fill := true
		for x := 0; x < m.width; x++ {
			if !m.cells[y][x] {
				fill = false
				break
			}
		}

		if fill {
			lines++

			for i := y - 1; i >= 0; i-- {
				for x := 0; x < m.width; x++ {
					m.cells[i+1][x] = m.cells[i][x]
					m.colors[i+1][x] = m.colors[i][x]
				}
			}
		}
	}
	return lines
}

func (m *matrix) GetCells() [][]Cell {
	return m.cells
}

func (m *matrix) GetColors() [][]Color {
	return m.colors
}
