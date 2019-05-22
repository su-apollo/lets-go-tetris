package client

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/pkg/game"
	"time"
)

const uiX = 4
const uiY = 4

// Client 구조체는 게임의 외형 정보를 저장한다.
type Client struct {
	Width, Height int
	CellSize      int
	Title         string
}

// Run 함수는 게임을 실행하는 메인 루프로 게임이 종료 될 때까지 블로킹 된다.
func (c *Client) Run() error {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var texture *sdl.Texture
	var err error

	width := int32((c.Width + uiX) * c.CellSize)
	height := int32(c.Height * c.CellSize)
	window, err = sdl.CreateWindow(c.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_OPENGL)

	if err != nil {
		return err
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		return err
	}
	defer renderer.Destroy()

	g := game.New(c.Width, c.Height)
	image, err := img.Load("./assets/gopher.png")
	if err != nil {
		return err
	}
	defer image.Free()

	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		return err
	}
	defer texture.Destroy()

	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Clear()
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
		}

		now := time.Now()
		delta := now.Sub(front)
		front = now

		g.Update(delta.Nanoseconds())

		renderer.Clear()
		renderer.SetDrawColor(0, 0, 0, 0xff)
		renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: width, H: height})

		drawBoard(renderer, g.Board, c.CellSize)
		drawBlock(renderer, g.CurrBlock, c.CellSize, 0)
		drawBlock(renderer, g.NextBlock, c.CellSize, c.Width)
		drawBlock(renderer, g.GetGhostBlock(), c.CellSize, 0)

		if g.State == game.Paused {
			w := int32((c.Width + uiX) * c.CellSize)
			h := int32(c.Height * c.CellSize)
			src := sdl.Rect{W: 172, H: 230}
			dst := sdl.Rect{X: 0, Y: 0, W: w, H: h}
			center := sdl.Point{
				X: dst.W / 2,
				Y: dst.H / 2,
			}
			renderer.CopyEx(texture, &src, &dst, 0, &center, 0)
		} else {
			x := int32((c.Width+uiX)*c.CellSize - 86)
			y := int32(c.Height*c.CellSize - 115)
			src := sdl.Rect{W: 172, H: 230}
			dst := sdl.Rect{X: x, Y: y, W: 86, H: 115}
			center := sdl.Point{
				X: dst.W / 2,
				Y: dst.H / 2,
			}
			renderer.CopyEx(texture, &src, &dst, 0, &center, 0)
		}

		renderer.Present()
	}

	return nil
}

func drawBoard(renderer *sdl.Renderer, board game.Board, size int) {
	for y, line := range board.GetCells() {
		for x := range line {
			color := board.GetColor(x, y)
			renderer.SetDrawColor(color.R, color.G, color.B, color.A)
			renderer.FillRect(&sdl.Rect{
				X: int32(x * size),
				Y: int32(y * size),
				W: int32(size),
				H: int32(size),
			})
		}
	}
}

func drawBlock(renderer *sdl.Renderer, block game.Block, size int, offsetX int) {
	posX, posY := block.GetPosition()
	color := block.GetColor()
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	posX += offsetX

	for y, line := range block.GetCells() {
		for x, cell := range line {
			if cell {
				cx := posX + x
				cy := posY + y

				renderer.FillRect(&sdl.Rect{
					X: int32(cx * size),
					Y: int32(cy * size),
					W: int32(size),
					H: int32(size),
				})
			}
		}
	}
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
