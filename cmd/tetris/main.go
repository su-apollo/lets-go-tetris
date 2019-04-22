package main

import (
	"lets-go-tetris/game"
	"lets-go-tetris/option"
	"lets-go-tetris/render/sdl"
)

func main() {
	opt := option.Opt{X: 11, Y: 23, CellSize: 20, Title: "Lets go"}
	r, err := sdl.NewSDLWrapper(opt)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	g := game.New(opt, r)
	g.Run()
}
