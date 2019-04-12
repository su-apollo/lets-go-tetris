package game

import (
	"lets-go-tetris/render"
	"math/rand"
)

const (
	shapeX = 4
	shapeY = 4
)

const S = `
0110
1200
0000
0000
`

const Z = `
1100
0210
0000
0000
`

const T = `
0100
1210
0000
0000
`

const I = `
0100
0200
0100
0100
`
const O = `
1100
1100
0000
0000
`

const L = `
0100
0200
0110
0000
`

const J = `
0100
0200
1100
0000
`

var shapes = []string{
	S,
	Z,
	T,
	I,
	O,
	L,
	J,
}

var colors = []uint32{
	0xffabf200,
	0xffff0000,
	0xffff00dd,
	0xff00d8ff,
	0xffffe400,
	0xffffbb00,
	0xff0100ff,
}

type cell bool

type mino struct {
	x, y  int32
	cells []cell
	color uint32
}

func random(seed int64) uint32 {
	rand.Seed(seed)
	return rand.Uint32() % uint32(len(shapes))
}

func NewRandomMino(seed int64) *mino {
	i := random(seed)
	b := &mino{color: colors[i]}
	b.init(shapes[i])
	return b
}

func (m *mino) RenderInfo() []render.Info {
	var infos []render.Info

	var x, y int32 = 0, 0
	for _, cell := range m.cells {
		if cell {
			infos = append(infos, render.Info{
				PosX: m.x + x, PosY: m.y + y, Color: m.color,
			})
		}
		x++
		if x%shapeX == 0 {
			x = 0
			y++
		}
	}

	return infos
}

func (m *mino) init(shape string) {
	m.cells = make([]cell, shapeX*shapeY)
	i := 0
	for _, c := range shape {
		switch c {
		case '2':
			fallthrough
		case '1':
			m.cells[i] = true
			fallthrough
		case '0':
			i++
		}
	}
}
