package game

var tileColors = []uint32{
	0xff353535,
	0xff5D5D5D,
}

type matrix struct {
	width, height int
	cells         []cell
	colors        []uint32
}

func (m *matrix) reset() {
	m.cells = make([]cell, m.width*m.height)
	m.colors = make([]uint32, m.width*m.height)
}

func (m *matrix) collide(b Block) bool {
	var x, y = 0, 0
	for _, cell := range b.GetCells() {
		if cell {
			cx, cy := b.GetPosition()
			cx += x
			cy += y

			if cx < 0 || m.width <= cx || cy < 0 || m.height <= cy {
				return true
			}

			if m.cells[cy*m.width+cx] {
				return true
			}
		}

		x++
		if x%shapeX == 0 {
			x = 0
			y++
		}
	}
	return false
}

func (m *matrix) merge(b Block) {
	var x, y = 0, 0
	for _, cell := range b.GetCells() {
		if cell {
			cx, cy := b.GetPosition()
			cx += x
			cy += y

			if 0 <= cx && cx < m.width && 0 <= cy && cy < m.height {
				m.cells[cy*m.width+cx] = true
				m.colors[cy*m.width+cx] = b.GetColor()
			}
		}

		x++
		if x%shapeX == 0 {
			x = 0
			y++
		}
	}
}

func (m *matrix) removeLines() int {
	lines := 0
	for y := 0; y < m.height; y++ {
		fill := true
		for x := 0; x < m.width; x++ {
			if !m.cells[y*m.width+x] {
				fill = false
				break
			}
		}

		if fill {
			lines++

			for i := y - 1; i >= 0; i-- {
				for x := 0; x < m.width; x++ {
					offset := i*m.width + x
					m.cells[offset+m.width] = m.cells[offset]
					m.colors[offset+m.width] = m.colors[offset]
				}
			}
		}
	}
	return lines
}

func (m *matrix) GetCells() []cell {
	return m.cells
}

func (m *matrix) GetColors() []uint32 {
	return m.colors
}
