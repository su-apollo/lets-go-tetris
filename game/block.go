package game

import (
	"fmt"
	"math/rand"
)

const startX = 3
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

type cell bool

type block struct {
	x, y  int
	cells []cell
}

func randomShape() string {
	i := rand.Uint32() % uint32(len(shapes))
	return shapes[i]
}

func NewRandomBlock() *block {
	b := &block{x: startX}
	b.init(randomShape())
	return b
}

func (b *block) init(shape string) {
	b.cells = make([]cell, shapeX*shapeY)
	i := 0
	for _, c := range shape {
		switch c {
		case '2':
			fallthrough
		case '1':
			b.cells[i] = true
			fallthrough
		case '0':
			i++
		}
	}
}

func (b *block) draw() {
	i := 0
	for _, cell := range b.cells {
		if i%shapeX == 0 {
			fmt.Print("\n")
		}

		if cell {
			fmt.Printf("██")
		} else {
			fmt.Printf("  ")
		}

		i++
	}
}
