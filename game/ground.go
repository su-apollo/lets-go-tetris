package game

import (
	"lets-go-tetris/render"
)

var tileColors = []uint32{
	0xff353535,
	0xff5D5D5D,
}

type ground struct {
	x, y   int32
	cells  []cell
	colors []uint32
}

func (g *ground) RenderInfo() []render.Info {
	var infos []render.Info

	var x, y int32 = 0, 0
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
		if x%g.x == 0 {
			x = 0
			y++
		}
	}

	return infos
}

func (g *ground) reset() {
	g.cells = make([]cell, g.x*g.y)
	g.colors = make([]uint32, g.x*g.y)
}

func (g *ground) step(m *mino) bool {
	m.y++
	if !g.collide(m) {
		return false
	}

	m.y--
	g.merge(m)

	return true
}

func (g *ground) collide(m *mino) bool {
	var x, y int32 = 0, 0
	for _, cell := range m.currentCells() {
		if cell {
			cx := m.x + x
			cy := m.y + y

			if cx < 0 || g.x <= cx || cy < 0 || g.y <= cy {
				return true
			}

			if g.cells[cy*g.x+cx] {
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

func (g *ground) merge(m *mino) {
	var x, y int32 = 0, 0
	for _, cell := range m.currentCells() {
		if cell {
			cx := m.x + x
			cy := m.y + y

			if 0 <= cx && cx < g.x && 0 <= cy && cy < g.y {
				g.cells[cy*g.x+cx] = true
				g.colors[cy*g.x+cx] = m.color
			}
		}

		x++
		if x%shapeX == 0 {
			x = 0
			y++
		}
	}
}

func (g *ground) tetris() int {
	return 0
}
