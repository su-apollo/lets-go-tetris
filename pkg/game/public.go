package game

// State is game state
type State int

// Playing
// ||     \\
// Paused  Over
const (
	Playing State = iota
	Paused
	Over
)


// Shape types of the block.
type Shape int

// https://tetris.wiki/Tetromino
const (
	I Shape = 0 + iota
	J
	L
	O
	S
	T
	Z
)


// It's the most basic unit in tetris.
type Cell bool

const (
	o Cell = true
	x Cell = false
)

// Abstract tetris block
type Block interface {
	Cells() [][]Cell
	Position() (int, int)
	Shape() Shape
}

// Abstract tetris board
type Board interface {
	Cells() [][]Cell
	CellShape(int, int) Shape
	Collide(Block) bool
}

// Abstract tetris game
type Game interface {
	State() State
	NowBlock() Block
	NextBlock() Block
	KeepBlock() Block
	GhostBlock() Block
	Board() Board

	HandleKey(msg Msg)
	Update(delta int64)
}

// Create new tetris game
func New(width int, height int) Game {
	t := &tetris{
		mat: &matrix{width, height, nil, nil},
	}
	t.reset()
	return t
}