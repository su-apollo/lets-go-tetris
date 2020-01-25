package main

import (
	"fmt"
	"lets-go-tetris/pkg/game-client"
	"os"
)

func main() {
	c := &game_client.Client{Width: 450, Height: 720, Title: "Let's go tetris!"}
	err := c.Run()

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
