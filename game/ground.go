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

func (g *ground) draw(s *sdl.Surface, c int) {
	var x, y = 0, 0

	for i, cell := range g.cells {
		if cell {

		} else {
			r := sdl.Rect{X: int32(x*c), Y: int32(y*c), W: int32(c), H: int32(c)}
			_ = s.FillRect(&r, tileColors[i%len(tileColors)])
		}

		x++
		if x%g.x == 0 {
			x = 0
			y++
		}
	}
}
