package main

import (
	"fmt"
	"math/rand"
)

const X = 11
const Y = 23

const X_Start = 3
const X_Shape = 4

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

var Shapes = [...]string{
	S,
	Z,
	T,
	I,
	O,
	L,
	J,
}

type Block struct {
	X int
	Y int
	Cells [16]Cell
}

type Cell struct {
	Filled bool
}

// todo :
func NewBlock(shape string) *Block {
	block := new(Block)
	block.X = X_Start

	i := 0
	for _, c := range shape {
		switch c {
		case '2':
			fallthrough
		case '1':
			block.Cells[i].Filled = true
			fallthrough
		case '0':
			i++
		}
	}

	return block
}

func RandomShape() string {
	i := rand.Uint32() % uint32(len(Shapes))
	return Shapes[i]
}

func NewRandomBlock() *Block {
	block := NewBlock(RandomShape())
	return block
}

func main() {
	block := NewRandomBlock()
	i := 0
	for _, c := range block.Cells {
		if i % X_Shape == 0 {
			fmt.Print("\n")
		}

		if c.Filled {
			fmt.Printf("██")
		} else {
			fmt.Printf("  ")
		}

		i++
	}
}