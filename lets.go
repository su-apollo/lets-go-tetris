package main

import (
	"fmt"
	"math/rand"
)

const X = 11
const Y = 23

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
const B = `
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
	B,
	L,
	J,
}

type Block struct {
	X int
	Y int
	Shape string
}

// todo :
func NewBlock() *Block {
	block := new(Block)
	return block
}

func RandomBlock(block *Block) *Block {
	i := rand.Uint32() % uint32(len(Shapes))
	block.Shape = Shapes[i]

	return block
}

func NewRandomBlock() *Block {
	block := RandomBlock(NewBlock())

	return block
}

func main() {
	block := NewRandomBlock()
	fmt.Println(block.Shape)
}