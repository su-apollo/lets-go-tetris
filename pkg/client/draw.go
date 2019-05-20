package client

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/pkg/game"
)

type draw struct {
	renderer *sdl.Renderer
	image    *sdl.Surface
	texture  *sdl.Texture
}

func (d *draw) init(window *sdl.Window) error {
	var err error

	d.renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		return err
	}

	d.image, err = img.Load("./assets/gopher/dancing_gopher.png")
	if err != nil {
		return err
	}

	d.texture, err = d.renderer.CreateTextureFromSurface(d.image)
	if err != nil {
		return err
	}

	err = d.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		fmt.Println(err)
	}

	err = d.renderer.Clear()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (d *draw) destroy() {
	d.renderer.Destroy()
	d.image.Free()
	d.texture.Destroy()
}

func (d *draw) clear(width int32, height int32) {
	err := d.renderer.Clear()
	if err != nil {
		fmt.Println(err)
	}

	err = d.renderer.SetDrawColor(0, 0, 0, 0xff)
	if err != nil {
		fmt.Println(err)
	}

	err = d.renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: width, H: height})
	if err != nil {
		fmt.Println(err)
	}
}

func (d *draw) swap() {
	d.renderer.Present()
}

func (d *draw) drawGame(game game.Game, size int) {
	d.drawBoard(game.Board(), size)
	d.drawBoard(game.Board(), size)
	d.drawNow(game, size)
	d.drawGhost(game, size)
}

func (d *draw) drawUI(g game.Game, c *Client) {
	d.drawNext(g, c)

	src := sdl.Rect{X: 192, Y: 192, W: 384, H: 384}
	dst := sdl.Rect{X: gopherX, Y: gopherY, W: 192, H: 192}
	center := sdl.Point{
		X: dst.W / 2,
		Y: dst.H / 2,
	}
	err := d.renderer.CopyEx(d.texture, &src, &dst, 0, &center, 0)
	if err != nil {
		fmt.Println(err)
	}
}

func (d *draw) drawBoard(board game.Board, size int) {
	for y, line := range board.Cells() {
		for x, cell := range line {
			var color Color
			if cell {
				color = colors[board.CellShape(x, y)]
			} else {
				color = *tileColor
			}

			err := d.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
			if err != nil {
				fmt.Println(err)
			}

			err = d.renderer.FillRect(&sdl.Rect{
				X: int32(x * size),
				Y: int32(y * size),
				W: int32(size),
				H: int32(size),
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (d *draw) drawNow(g game.Game, size int) {
	color := colors[g.NowBlock().Shape()]
	d.drawBlock(g.NowBlock(), color, size, 0)
}

func (d *draw) drawGhost(g game.Game, size int) {
	color := colors[g.GhostBlock().Shape()]
	color.R &= ghostMask.R
	color.G &= ghostMask.G
	color.B &= ghostMask.B
	color.A &= ghostMask.A
	d.drawBlock(g.GhostBlock(), color, size, 0)
}

func (d *draw) drawNext(g game.Game, c *Client) {
	color := colors[g.NextBlock().Shape()]
	size := cellSize

	err := d.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	if err != nil {
		fmt.Println(err)
	}

	for y, line := range g.NextBlock().Cells() {
		for x, cell := range line {
			if cell {
				err = d.renderer.FillRect(&sdl.Rect{
					X: int32(x*size + nextX),
					Y: int32(y*size + nextY),
					W: int32(size),
					H: int32(size),
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func (d *draw) drawBlock(block game.Block, color Color, size int, offsetX int) {
	posX, posY := block.Position()

	err := d.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	if err != nil {
		fmt.Println(err)
	}

	for y, line := range block.Cells() {
		for x, cell := range line {
			if cell {
				cx := posX + x
				cy := posY + y

				err = d.renderer.FillRect(&sdl.Rect{
					X: int32(cx * size),
					Y: int32(cy * size),
					W: int32(size),
					H: int32(size),
				})
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
