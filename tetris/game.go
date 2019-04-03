package tetris

const GroundX = 11
const GroundY = 23

const (
	Playing = iota
	Paused
	GameOver
)

type Game struct {
	State int
	Now   *Block
	Next  *Block
	Back  *Ground
}

func NewGame() *Game {
	game := new(Game)

	game.State = Playing
	game.Now = NewRandomBlock()
	game.Next = NewRandomBlock()
	game.Back = NewGround(GroundX, GroundY)

	return game
}

func (game *Game) Draw() {
	game.Back.Draw()
	game.Now.Draw()
}
