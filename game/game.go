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

const startX = 3

type Game struct {
	state State
	now   *mino
	next  *mino
	back  *ground

	render render.Renderer

	stepTimer int64
}

func New(opt option.Opt, r render.Renderer) *Game {
	g := &ground{opt.X, opt.Y, nil}
	g.reset()

	now := NewRandomMino(time.Now().UnixNano())
	now.x = startX

	next := NewRandomMino(time.Now().UnixNano() + 1)
	next.x = opt.X

	return &Game{
		state:  Playing,
		now:    now,
		next:   next,
		back:   g,
		render: r,
	}
}

func (game *Game) Run() {
	var info []render.Info
	front := time.Now()
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

		now := time.Now()
		delta := now.Sub(front)
		front = now

		game.update(delta.Nanoseconds())
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

func (game *Game) update(delta int64) {
	switch game.state {
	case Playing:
		game.updatePlaying(delta)
		break
	case Paused:
		game.updatePaused(delta)
		break
	case Over:
		game.updateGameOver(delta)
		break
	}
}

func (game *Game) updatePlaying(delta int64) {
	game.stepTimer += delta
	if game.stepTimer > game.speed() {
		if game.back.step(game.now) {
			_ = game.back.tetris()
			//todo : score

			game.now = game.next
			game.now.x = startX
			game.next = NewRandomMino(time.Now().UnixNano())
			game.next.x = game.back.x
		}
		game.stepTimer = 0
	}
}

func (game *Game) updatePaused(delta int64) {
	panic("Not implemented")
}

func (game *Game) updateGameOver(delta int64) {
	panic("Not implemented")
}

func (game *Game) speed() int64 {
	// todo : game level
	return 1000000000
}