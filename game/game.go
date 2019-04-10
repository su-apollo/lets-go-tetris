package game

import (
	"lets-go-tetris/event"
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

func New(opt option.Opt, r render.Renderer) *Game {
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
	var info []render.Info
	for {
		info = info[:0]

		info = append(info, game.back.RenderInfo()...)
		info = append(info, game.now.RenderInfo()...)
		info = append(info, game.next.RenderInfo()...)

		game.render.Render(info)

		keys, ok := game.render.Update()
		if !ok {
			break
		}

		for _, key := range keys {
			game.handleKey(key)
		}
	}
}

func (game *Game) handleKey(msg event.Msg) {
	switch game.state {
	case Playing:
		game.handleKeyPlaying(msg)
		break
	case Paused:
		game.handleKeyPaused(msg)
		break
	case Over:
		game.handleKeyGameOver(msg)
		break
	}
}

func (game *Game) handleKeyPlaying(msg event.Msg) {
	switch msg.Key {
	case event.Left:
	case event.Right:
	case event.Up:
	case event.Down:
	case event.Escape:
	case event.Pause:
	}
}

func (game *Game) handleKeyPaused(msg event.Msg) {
	panic("Not implemented")
}

func (game *Game) handleKeyGameOver(msg event.Msg) {
	panic("Not implemented")
}
