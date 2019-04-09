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
				PosX: int32(x), PosY: int32(y), Color: tileColors[i%len(tileColors)],
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
