package game

import (
	"lets-go-tetris/render"
)

var tileColors = []uint32{
	0xff353535,
	0xff5D5D5D,
}

type ground struct {
	width, height int
	cells         []cell
	colors        []uint32
}

func (g *ground) RenderInfo() []render.Info {
	var infos []render.Info

	var x, y = 0, 0
	for i, cell := range g.cells {
		if !cell {
			infos = append(infos, &render.InfoImpl{
				PosX: x, PosY: y, Color: tileColors[i%len(tileColors)],
			})
		} else {
			infos = append(infos, &render.InfoImpl{
				PosX: x, PosY: y, Color: g.colors[i],
			})
		}
		x++
		if x%g.width == 0 {
			x = 0
			y++
		}
	}

	return infos
}

func (g *ground) reset() {
	g.cells = make([]cell, g.width*g.height)
	g.colors = make([]uint32, g.width*g.height)
}

func (g *ground) step(m *tetromino) bool {
	m.y++
	if !g.collide(m) {
		return false
	}

	m.y--
	g.merge(m)

	return true
}

func (g *ground) collide(m *tetromino) bool {
	var x, y = 0, 0
	for _, cell := range m.currentCells() {
		if cell {
			cx := m.x + x
			cy := m.y + y

			if cx < 0 || g.width <= cx || cy < 0 || g.height <= cy {
				return true
			}

			if g.cells[cy*g.width+cx] {
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

func (g *ground) merge(m *tetromino) {
	var x, y = 0, 0
	for _, cell := range m.currentCells() {
		if cell {
			cx := m.x + x
			cy := m.y + y

			if 0 <= cx && cx < g.width && 0 <= cy && cy < g.height {
				g.cells[cy*g.width+cx] = true
				g.colors[cy*g.width+cx] = m.color
			}
		}

		x++
		if x%shapeX == 0 {
			x = 0
			y++
		}
	}
}

func (g *ground) removeLines() int {
	lines := 0
	for y := 0; y < g.height; y++ {
		fill := true
		for x := 0; x < g.width; x++ {
			if !g.cells[y*g.width+x] {
				fill = false
				break
			}
		}

		if fill {
			lines++

			for i := y - 1; i >= 0; i-- {
				for x := 0; x < g.width; x++ {
					offset := i*g.width + x
					g.cells[offset+g.width] = g.cells[offset]
					g.colors[offset+g.width] = g.colors[offset]
				}
			}
		}
	}
	return lines
}
