package main

import (
	"lets-go-tetris/pkg/game-server"
)

func main() {
	s := &game_server.Server{}
	s.Run()
}
