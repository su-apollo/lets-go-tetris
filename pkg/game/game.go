package game

import (
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
	now   *tetromino
	next  *tetromino
	mat   *matrix

	stepTimer int64
}

func (g *Game) HandleKey(msg Msg) {
	switch g.state {
	case Playing:
		g.handleKeyPlaying(msg)
		break
	case Paused:
		g.handleKeyPaused(msg)
		break
	case Over:
		g.handleKeyGameOver(msg)
		break
	}
}

func (g *Game) Update(delta int64) {
	switch g.state {
	case Playing:
		g.updatePlaying(delta)
		break
	case Paused:
		g.updatePaused(delta)
		break
	case Over:
		g.updateGameOver(delta)
		break
	}
}

func New(width int, height int) *Game {
	m := &matrix{width, height, nil, nil}
	m.reset()

	rand.Seed(time.Now().UnixNano())
	now := randomTetromino()
	now.x = startX

	next := randomTetromino()
	next.x = width

	return &Game{
		state: Playing,
		now:   now,
		next:  next,
		mat:   m,
	}
}

func (g *Game) handleKeyPlaying(msg Msg) {
	switch msg.Key {
	case Left:
		g.now.x--
		if g.mat.collide(g.now) {
			g.now.x++
		}
	case Right:
		g.now.x++
		if g.mat.collide(g.now) {
			g.now.x--
		}
	case Down:
		g.now.y++
		if g.mat.collide(g.now) {
			g.now.y--
			g.nextStep()
		}
	case ClockWise:
		r := g.now.rotateClockWise()
		if !g.now.wallKick(g.mat, r) {
			g.now.rotateCounterClockWise()
		}
	case CounterClockWise:
		r := g.now.rotateCounterClockWise()
		if !g.now.wallKick(g.mat, r) {
			g.now.rotateClockWise()
		}
	case Drop:
		drop := true
		for drop {
			g.now.y++
			drop = !g.mat.collide(g.now)
		}
		g.now.y--
		g.nextStep()
	case Escape:
		g.state = Over
	case Pause:
		g.state = Paused
	}
}

func (g *Game) handleKeyPaused(msg Msg) {
	switch msg.Key {
	case Pause:
		g.state = Playing
	}
}

func (g *Game) handleKeyGameOver(msg Msg) {
}

func (g *Game) updatePlaying(delta int64) {
	g.stepTimer += delta
	if g.stepTimer > g.speed() {
		if g.step(g.mat, g.now) {
			_ = g.mat.removeLines()
			//todo : score

			g.now = g.next
			g.now.x = startX
			g.next = randomTetromino()
			g.next.x = g.mat.width

			if g.mat.collide(g.now) {
				g.state = Over
			}
		}
		g.stepTimer = 0
	}
}

func (g *Game) updatePaused(delta int64) {
}

func (g *Game) updateGameOver(delta int64) {
	g.mat.reset()

	rand.Seed(time.Now().UnixNano())
	g.now = randomTetromino()
	g.now.x = startX

	g.next = randomTetromino()
	g.next.x = g.mat.width

	g.state = Playing
}

func (g *Game) speed() int64 {
	// todo : g level
	return 1000000000
}

func (g *Game) nextStep() {
	g.stepTimer += g.speed()
}

func (g *Game) step(m *matrix, t *tetromino) bool {
	t.y++
	if !m.collide(t) {
		return false
	}

	t.y--
	m.merge(t)

	return true
}

func (g *Game) GetNowBlock() Block {
	return g.now
}

func (g *Game) GetNextBlock() Block {
	return g.next
}

func (g *Game) GetGhostBlock() Block {
	ghost := &ghost{}
	ghost.init(g.mat, g.now)
	return ghost
}

func (g *Game) GetBoard() Board {
	return g.mat
}
