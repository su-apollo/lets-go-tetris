package client

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/pkg/game"
)

type draw struct {
	renderer *sdl.Renderer
	texture *sdl.Texture
}

func (d *draw) init(window *sdl.Window) error {
	var err error

	d.renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		return err
	}
	defer d.renderer.Destroy()

	image, err := img.Load("./assets/gopher.png")
	if err != nil {
		return err
	}
	defer image.Free()

	d.texture, err = d.renderer.CreateTextureFromSurface(image)
	if err != nil {
		return err
	}
	defer d.texture.Destroy()

	d.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	d.renderer.Clear()

	return nil
}

func (d *draw) clear(width int32, height int32) {
	d.renderer.Clear()
	d.renderer.SetDrawColor(0, 0, 0, 0xff)
	d.renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: width, H: height})
}

func (d *draw) swap() {
	d.renderer.Present()
}

func (d *draw) drawGame(game game.Game, size int) {
	d.drawBoard(game.Board(), size)
	d.drawBoard(game.Board(), size)
	d.drawBlock(game.NowBlock(), size, 0)
	d.drawBlock(game.GhostBlock(), size, 0)
}

func (d *draw) drawUI(g game.Game, c *Client) {
	d.drawBlock(g.NextBlock(), c.CellSize, c.Width)

	if g.State() == game.Paused {
		w := int32((c.Width + uiX) * c.CellSize)
		h := int32(c.Height * c.CellSize)
		src := sdl.Rect{W: 172, H: 230}
		dst := sdl.Rect{X: 0, Y: 0, W: w, H: h}
		center := sdl.Point{
			X: dst.W / 2,
			Y: dst.H / 2,
		}
		d.renderer.CopyEx(d.texture, &src, &dst, 0, &center, 0)
	} else {
		x := int32((c.Width+uiX)*c.CellSize - 86)
		y := int32(c.Height*c.CellSize - 115)
		src := sdl.Rect{W: 172, H: 230}
		dst := sdl.Rect{X: x, Y: y, W: 86, H: 115}
		center := sdl.Point{
			X: dst.W / 2,
			Y: dst.H / 2,
		}
		d.renderer.CopyEx(d.texture, &src, &dst, 0, &center, 0)
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

			d.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
			d.renderer.FillRect(&sdl.Rect{
				X: int32(x * size),
				Y: int32(y * size),
				W: int32(size),
				H: int32(size),
			})
		}
	}
}

func (d *draw) drawBlock(block game.Block, size int, offsetX int) {
	posX, posY := block.Position()
	color := colors[block.Shape()]
	d.renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	posX += offsetX

	for y, line := range block.Cells() {
		for x, cell := range line {
			if cell {
				cx := posX + x
				cy := posY + y

				d.renderer.FillRect(&sdl.Rect{
					X: int32(cx * size),
					Y: int32(cy * size),
					W: int32(size),
					H: int32(size),
				})
			}
		}
	}
}
