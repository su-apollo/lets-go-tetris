package game

import (
	"lets-go-tetris/interfaces/renderer"
	"lets-go-tetris/render"
	"math/rand"
)

const (
	shapeX      = 4
	shapeY      = 4
	rotationMax = 4
)

type Shape int

const (
	I Shape = 0 + iota
	J
	L
	O
	S
	T
	Z
)

const i0 = `
0000
1111
0000
0000
`

const i1 = `
0010
0010
0010
0010
`

const i2 = `
0000
0000
1111
0000
`

const i3 = `
0100
0100
0100
0100
`

const j0 = `
1000
1110
0000
0000
`

const j1 = `
0110
0100
0100
0000
`

const j2 = `
0000
1110
0010
0000
`

const j3 = `
0100
0100
1100
0000
`

const s0 = `
0110
1100
0000
0000
`

const s1 = `
0100
0110
0010
0000
`

const s2 = `
0000
0110
1100
0000
`

const s3 = `
1000
1100
0100
0000
`

const z0 = `
1100
0110
0000
0000
`

const z1 = `
0010
0110
0100
0000
`

const z2 = `
0000
1100
0110
0000
`
const z3 = `
0100
1100
1000
0000
`

const t0 = `
0100
1110
0000
0000
`

const t1 = `
0100
0110
0100
0000
`
const t2 = `
0000
1110
0100
0000
`

const t3 = `
0100
1100
0100
0000
`

const o0 = `
0110
0110
0000
0000
`

const l0 = `
0010
1110
0000
0000
`

const l1 = `
0100
0100
0110
0000
`

const l2 = `
0000
1110
1000
0000
`

const l3 = `
1100
0100
0100
0000
`

var shapes = [][]string{
	{i0, i1, i2, i3},
	{j0, j1, j2, j3},
	{l0, l1, l2, l3},
	{o0, o0, o0, o0},
	{s0, s1, s2, s3},
	{t0, t1, t2, t3},
	{z0, z1, z2, z3},
}

var colors = []uint32{
	0xff00d8ff,
	0xff0100ff,
	0xffffbb00,
	0xffffe400,
	0xffabf200,
	0xffff00dd,
	0xffff0000,
}

type cell bool

type mino struct {
	x, y     int32
	cells    [][]cell
	color    uint32
	rotation int
}

func randomMino() *mino {
	i := rand.Intn(len(shapes) - 1)
	m := &mino{color: colors[i]}
	m.init(shapes[i])
	return m
}

func newMino(s Shape) *mino {
	m := &mino{color: colors[s]}
	m.init(shapes[s])
	return m
}

func (m *mino) RenderInfo() []renderer.Info {
	var infos []renderer.Info

	var x, y int32 = 0, 0
	for _, cell := range m.currentCells() {
		if cell {
			infos = append(infos, &render.InfoImpl{
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

func (m *mino) init(rotationShapes []string) {
	m.cells = make([][]cell, rotationMax)
	for i := range m.cells {
		m.cells[i] = make([]cell, shapeX*shapeY)
	}

	for r, shape := range rotationShapes {
		i := 0
		for _, c := range shape {
			switch c {
			case '1':
				m.cells[r][i] = true
				fallthrough
			case '0':
				i++
			}
		}
	}
}

func (m *mino) rotate(r int) {
	m.rotation = (r%rotationMax + rotationMax) % rotationMax
}

func (m *mino) currentCells() []cell {
	return m.cells[m.rotation]
}
