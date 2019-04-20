package game

import (
	"lets-go-tetris/event"
	"lets-go-tetris/interfaces/renderer"
	"lets-go-tetris/option"
	"math/rand"
	"time"
)

// State 타입은 게임의 상태를 나타낸다.
type State int

// Playing 		게임 진행 중
// Paused 		일시 정지
// Over			게임 종료
const (
	Playing State = iota
	Paused
	Over
)

const startX = 3

// Game 구조체는 테트리스의 전반 로직을 담당하는 자료구조다.
type Game struct {
	state State
	now   *mino
	next  *mino
	back  *ground

	render renderer.Renderer

	stepTimer int64
}

// New 함수는 게임 실행 옵션과 화면에 출력을 담당할 렌더러를 전달 받고 테트리스 로직을 담은 Game 자료구조를 반환한다.
func New(opt option.Opt, r renderer.Renderer) *Game {
	g := &ground{opt.X, opt.Y, nil, nil}
	g.reset()

	rand.Seed(time.Now().UnixNano())
	now := randomMino()
	now.x = startX

	next := randomMino()
	next.x = opt.X

	return &Game{
		state:  Playing,
		now:    now,
		next:   next,
		back:   g,
		render: r,
	}
}

// Run 함수는 블로킹 된 상태로 게임을 실행한다.
func (game *Game) Run() {
	var info []renderer.Info
	front := time.Now()
	for {
		info = info[:0]

		info = append(info, game.back.RenderInfo()...)
		info = append(info, game.now.RenderInfo()...)
		info = append(info, game.next.RenderInfo()...)

		if err := game.render.Render(info); err != nil {
			break
		}

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
			game.next = randomMino()
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
