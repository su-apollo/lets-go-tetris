package game

const ghostColor = 0x77ffffff

type ghost struct {
	cells    []cell
	x, y     int
}

func (g *ghost) init(gr *ground, mino *tetromino) {
}