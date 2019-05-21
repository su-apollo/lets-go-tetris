package client

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/pkg/game"
	"log"
	"net"
	"time"
)

//
type Client struct {
	Width, Height int
	Title         string

	draw *draw
}

//
func (c *Client) Run() error {
	conn, err := net.Dial("tcp", ":6000")
	if err != nil {
		log.Println(err)
	}

	if conn != nil {
		go func(conn net.Conn) {
			//writer := bufio.NewWriter(conn)
			for {
				s := "ping\n"
				//writer.WriteString(s)

				_, err := conn.Write([]byte(s))
				if err != nil {
					break
				}

				time.Sleep(time.Duration(3) * time.Second)
				fmt.Println("send ping")
			}
		}(conn)
	}

	width := int32(c.Width)
	height := int32(c.Height)
	window, err := sdl.CreateWindow(c.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_OPENGL)

	if err != nil {
		return err
	}
	defer window.Destroy()

	c.draw = &draw{}
	err = c.draw.init(window)

	if err != nil {
		return err
	}
	defer c.draw.destroy()

	g := game.New(game.BoardWidth, game.BoardHeight)
	d := game.New(game.BoardWidth, game.BoardHeight)

	front := time.Now()
	running := true

	for running {
		var keys []game.Msg
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if e.GetType() == sdl.KEYDOWN {
					if msg, ok := sdlKeyCodeToEvent(t.Keysym.Sym); ok {
						keys = append(keys, msg)
					}
				}
			}
		}

		for _, key := range keys {
			g.HandleKey(key)
			d.HandleKey(key)
		}

		now := time.Now()
		delta := now.Sub(front)
		front = now

		g.Update(delta.Nanoseconds())
		d.Update(delta.Nanoseconds())

		c.draw.clear(width, height)

		c.draw.drawGame(g, cellSize)
		c.draw.drawUI(g, c)
		c.draw.drawGopher(delta.Nanoseconds())

		c.draw.swap()
	}

	return nil
}

func sdlKeyCodeToEvent(k sdl.Keycode) (game.Msg, bool) {
	switch k {
	case sdl.K_LEFT, sdl.K_a, sdl.K_j:
		return game.Msg{Key: game.Left}, true
	case sdl.K_RIGHT, sdl.K_d, sdl.K_l:
		return game.Msg{Key: game.Right}, true
	case sdl.K_UP, sdl.K_w, sdl.K_i:
		return game.Msg{Key: game.ClockWise}, true
	case sdl.K_DOWN, sdl.K_s, sdl.K_k:
		return game.Msg{Key: game.Down}, true
	case sdl.K_SPACE:
		return game.Msg{Key: game.Drop}, true
	case sdl.K_ESCAPE:
		return game.Msg{Key: game.Escape}, true
	case sdl.K_p:
		return game.Msg{Key: game.Pause}, true
	default:
		return game.Msg{Key: game.Nop}, false
	}
}
