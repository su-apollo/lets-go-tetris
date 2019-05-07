package game

import "lets-go-tetris/render"

const ghostColor = 0x77ffffff

type ghost struct {
	cells    []cell
	x, y     int
	color    uint32
}

func (ghost *ghost) init(g *ground, m *tetromino) {
	ghost.x = m.x
	ghost.y = m.y
	ghost.cells = m.getCells()
	ghost.color = m.getColor() & ghostColor

	drop := true
	for drop {
		ghost.y++
		drop = !g.collide(ghost)
	}
	ghost.y--
}

func (ghost *ghost) getCells() []cell {
	return ghost.cells
}

func (ghost *ghost) getPosition() (int, int) {
	return ghost.x, ghost.y
}

func (ghost *ghost) getColor() uint32 {
	return ghost.color
}

func (ghost *ghost) RenderInfo() []render.Info {
	var infos []render.Info

	var x, y = 0, 0
	for _, cell := range ghost.getCells() {
		if cell {
			infos = append(infos, &render.InfoImpl{
				PosX: ghost.x + x, PosY: ghost.y + y, Color: ghost.color,
			})
		}
		x++
		if x%shapeX == 0 {
			x = 0
			y++
		}
	}

	return infos
}