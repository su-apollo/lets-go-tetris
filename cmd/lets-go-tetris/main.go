package main

import (
	"fmt"
	"lets-go-tetris/pkg/client"
	"os"
)

func main() {
	c := &client.Client{Width: 11, Height: 23, CellSize: 20, Title: "Let's go tetris!"}
	err := c.Run()

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
