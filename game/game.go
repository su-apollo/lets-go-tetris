package game

type State int

const (
	Playing State = iota
	Paused
	GameOver
)

type Option struct {
	X, Y int
}

type Game struct {
	State State
	Now   *block
	Next  *block
	Back  *ground
}

func New(opt Option) *Game {
	g := &ground{opt.X, opt.Y, nil}
	g.reset()
	return &Game{
		State: Playing,
		Now:   NewRandomBlock(),
		Next:  NewRandomBlock(),
		Back:  g,
	}
}

func (game *Game) Draw() {
	game.Back.draw()
	game.Now.draw()
}
