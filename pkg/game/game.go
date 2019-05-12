package game

import (
	"math/rand"
	"time"
)

const startX = 3

// Game 구조체는 테트리스의 전반 로직을 담당하는 자료구조다.
type Game struct {
	state State
	now   *tetromino
	next  *tetromino
	keep  *tetromino
	stack []Shape
	mat   *matrix

	stepTimer int64
}

func (g *Game) GetState() State {
	return g.state
}

func (g *Game) GetNowBlock() Block {
	return g.now
}

func (g *Game) GetNextBlock() Block {
	return g.next
}

func (g *Game) getKeepBlock() Block {
	return g.keep
}

func (g *Game) GetGhostBlock() Block {
	ghost := &ghost{}
	ghost.init(g.mat, g.now)
	return ghost
}

func (g *Game) GetBoard() Board {
	return g.mat
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
	g := &Game{
		mat: &matrix{width, height, nil, nil},
	}
	g.reset()
	return g
}

func (g *Game) setNowToNext() {
	g.now = g.next
	g.now.x = startX

	s := g.popQueue()
	g.next = newTetromino(s)

	if g.stack == nil {
		g.resetStack()
		g.shuffleStack()
	}
}

func (g *Game) reset() {
	g.mat.reset()
	g.resetStack()
	g.shuffleStack()

	s := g.popQueue()
	g.now = newTetromino(s)
	g.now.x = startX

	s = g.popQueue()
	g.next = newTetromino(s)

	g.state = Playing
}

func (g *Game) resetStack() {
	g.stack = []Shape{I, J, L, O, S, T, Z}
}

func (g *Game) shuffleStack() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(g.stack); n > 0; n-- {
		i := r.Intn(n)
		g.stack[n-1], g.stack[i] = g.stack[i], g.stack[n-1]
	}
}

func (g *Game) popQueue() Shape {
	n := len(g.stack)
	if n < 2 {
		v := g.stack[0]
		g.stack = nil
		return v
	} else {
		v := g.stack[n-1]
		g.stack = g.stack[:n-1]
		return v
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

			g.setNowToNext()

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
	g.reset()
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
