package main

import (
	"lets-go-tetris/game"
	"lets-go-tetris/option"
)

func main() {
	g := game.New(
		option.Opt{X: 11, Y: 23, CellSize: 20, Title: "Lets go"},
	)
	g.Run()
}
