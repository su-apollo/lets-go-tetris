package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/option"
	"lets-go-tetris/render"
	"time"
)

type State int

const (
	Playing State = iota
	Paused
	Over
)

type Game struct {
	state State
	now   *mino
	next  *mino
	back  *ground

	render render.Renderer
}

func New(opt option.Opt, r *render.SDLWrapper) *Game {
	g := &ground{opt.X, opt.Y, nil}
	g.reset()

	next := NewRandomMino(time.Now().UnixNano() + 1)
	next.state = Prepare
	next.offset = opt.X

	return &Game{
		state:  Playing,
		now:    NewRandomMino(time.Now().UnixNano()),
		next:   next,
		back:   g,
		render: r,
	}
}

func (game *Game) Run() {
	for {
		game.render.Render(game.back)
		game.render.Render(game.now)
		game.render.Render(game.next)

		keys, ok := game.render.Update()
		if !ok {
			break
		}

		for _, key := range keys {
			game.handleKey(key)
		}
	}
}

func (game *Game) handleKey(k sdl.Keycode) {
	switch game.state {
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
