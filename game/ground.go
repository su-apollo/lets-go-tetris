package game

import (
	"lets-go-tetris/render"
)

var tileColors = []uint32{
	0xffdbdbdb,
	0xffeaeaea,
}

type ground struct {
	x, y  int32
	cells []cell
}

func (g *ground) RenderInfo() []render.Info {
	var infos []render.Info

	var x, y int32 = 0, 0
	for i, cell := range g.cells {
		if !cell {
			infos = append(infos, render.Info{
				PosX: x, PosY: y, Color: tileColors[i%len(tileColors)],
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
	for _, cell := range m.cells {
		if cell {
			cx := m.x + x
			cy := m.y + y

			if cx < 0 || g.x <= cx || cy < 0 || g.y <= cy {
				return true
			}

			if g.cells[cy * g.x + cx] {
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
	for _, cell := range m.cells {
		if cell {
			cx := m.x + x
			cy := m.y + y

			g.cells[cy * g.x + cx] = true
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