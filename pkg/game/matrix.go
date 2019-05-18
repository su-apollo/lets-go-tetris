package game

type matrix struct {
	width, height int
	cells         [][]Cell
	shapes        [][]Shape
}

func (m matrix) Cells() [][]Cell {
	return m.cells
}

func (m matrix) CellShape(x int, y int) Shape {
	return m.shapes[y][x]
}

func (m matrix) Collide(b Block) bool {
	for y, cells := range b.Cells() {
		for x, cell := range cells {
			if cell {
				cx, cy := b.Position()
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

func newMatrix(w int, h int) *matrix {
	m := matrix{width: w, height: h}
	m.reset()
	return &m
}

func (m *matrix) reset() {
	m.cells = make([][]Cell, m.height)
	for i := range m.cells {
		m.cells[i] = make([]Cell, m.width)
	}

	m.shapes = make([][]Shape, m.height)
	for i := range m.shapes {
		m.shapes[i] = make([]Shape, m.width)
	}
}

func (m *matrix) merge(b Block) {
	for y, cells := range b.Cells() {
		for x, cell := range cells {
			if cell {
				cx, cy := b.Position()
				cx += x
				cy += y

				if 0 <= cx && cx < m.width && 0 <= cy && cy < m.height {
					m.cells[cy][cx] = true
					m.shapes[cy][cx] = b.Shape()
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
					m.shapes[i+1][x] = m.shapes[i][x]
				}
			}
		}
	}
	return lines
}
