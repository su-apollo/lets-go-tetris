package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/option"
	"time"
)

type State int

const (
	Playing State = iota
	Paused
	Over
)

type Game struct {
	State    State
	Now      *mino
	Next     *mino
	Back     *ground
	CellSize int
}

func New(opt option.Opt) *Game {
	g := &ground{opt.X, opt.Y, nil}
	g.reset()
	return &Game{
		State:    Playing,
		Now:      NewRandomMino(time.Now().UnixNano()),
		Next:     NewRandomMino(time.Now().UnixNano() + 1),
		Back:     g,
		CellSize: opt.CellSize,
	}
}

func (game *Game) Run() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	width := int32((game.Back.x + shapeX) * game.CellSize)
	height := int32(game.Back.y * game.CellSize)
	window, err := sdl.CreateWindow("lets go", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	running := true
	for running {
		game.draw(surface)
		window.UpdateSurface()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				if event.GetType() == sdl.KEYDOWN {
					game.handleKey(t.Keysym.Sym)
				}
				break
			}
		}
	}
}

func (game *Game) draw(s *sdl.Surface) {
	game.Back.draw(s, game.CellSize)
	game.Now.draw(s, game.CellSize, 0, 0)
	game.Next.draw(s, game.CellSize, game.Back.x, 0)
}

func (game *Game) handleKey(k sdl.Keycode) {
	switch game.State {
	case Playing:
		game.handleKeyPlaying(k)
		break
	case Paused:
		game.handleKeyPaused(k)
		break
	case Over:
		game.handleKeyGameOver(k)
		break
	}
}

func (game *Game) handleKeyPlaying(k sdl.Keycode) {
	switch k {
	case sdl.K_LEFT, sdl.K_a, sdl.K_j:

	case sdl.K_RIGHT, sdl.K_d, sdl.K_l:

	case sdl.K_UP, sdl.K_w, sdl.K_i:

	case sdl.K_DOWN, sdl.K_s, sdl.K_k, sdl.K_SPACE:

	case sdl.K_ESCAPE:

	case sdl.K_p:
	}
}

func (game *Game) handleKeyPaused(k sdl.Keycode) {
}

func (game *Game) handleKeyGameOver(k sdl.Keycode) {
}
