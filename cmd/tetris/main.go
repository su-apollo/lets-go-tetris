package main

import "lets-go-tetris/game"

func main() {
	g := game.New(
		game.Option{X: 11, Y: 23},
	)
	g.Draw()
}
