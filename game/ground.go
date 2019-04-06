package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

var tileColors = []uint32{
	0xffdbdbdb,
	0xffeaeaea,
}



type ground struct {
	x, y  int
	cells []cell
}

func (g *ground) reset() {
	g.cells = make([]cell, g.x*g.y)
}

func (g *ground) draw(s *sdl.Surface, cellSize int) {
	var x, y = 0, 0

	for i, cell := range g.cells {
		if cell {

		} else {
			r := sdl.Rect{X: int32(x*cellSize), Y: int32(y*cellSize), W: int32(cellSize), H: int32(cellSize)}
			_ = s.FillRect(&r, tileColors[i%len(tileColors)])
		}

		x++
		if x%g.x == 0 {
			x = 0
			y++
		}
	}
}
