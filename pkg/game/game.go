package game

import (
	"math/rand"
	"time"
)

const startX = 3

// Game 구조체는 테트리스의 전반 로직을 담당하는 자료구조다.
type Game struct {
	State     State
	CurrBlock *tetromino
	NextBlock *tetromino
	KeepBlock *tetromino
	stack     []Shape
	Board     *matrix

	stepTimer int64
}

// GetGhostBlock 메소드는 현재 블록의 미리보기 블록을 반환한다.
func (g *Game) GetGhostBlock() Block {
	ghost := &ghost{}
	ghost.init(g.Board, g.CurrBlock)
	return ghost
}

// HandleKey 메소드는 입력 받은 키보드 이벤트를 처리한다.
func (g *Game) HandleKey(msg Msg) {
	switch g.State {
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

// Update 메소드는 매 프레임을 갱신한다. 전달인자는 이전 프레임으로부터의 경과 나노초.
func (g *Game) Update(delta int64) {
	switch g.State {
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

// New 함수는 Game 자료구조의 생성자다.
func New(width int, height int) *Game {
	g := &Game{
		Board: &matrix{width, height, nil, nil},
	}
	g.reset()
	return g
}

func (g *Game) setNowToNext() {
	g.CurrBlock = g.NextBlock
	g.CurrBlock.x = startX

	s := g.popQueue()
	g.NextBlock = newTetromino(s)

	if g.stack == nil {
		g.resetStack()
		g.shuffleStack()
	}
}

func (g *Game) reset() {
	g.Board.reset()
	g.resetStack()
	g.shuffleStack()

	s := g.popQueue()
	g.CurrBlock = newTetromino(s)
	g.CurrBlock.x = startX

	s = g.popQueue()
	g.NextBlock = newTetromino(s)

	g.State = Playing
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
	}

	v := g.stack[n-1]
	g.stack = g.stack[:n-1]
	return v

}

func (g *Game) handleKeyPlaying(msg Msg) {
	switch msg.Key {
	case Left:
		g.CurrBlock.x--
		if g.Board.collide(g.CurrBlock) {
			g.CurrBlock.x++
		}
	case Right:
		g.CurrBlock.x++
		if g.Board.collide(g.CurrBlock) {
			g.CurrBlock.x--
		}
	case Down:
		g.CurrBlock.y++
		if g.Board.collide(g.CurrBlock) {
			g.CurrBlock.y--
			g.nextStep()
		}
	case ClockWise:
		r := g.CurrBlock.rotateClockWise()
		if !g.CurrBlock.wallKick(g.Board, r) {
			g.CurrBlock.rotateCounterClockWise()
		}
	case CounterClockWise:
		r := g.CurrBlock.rotateCounterClockWise()
		if !g.CurrBlock.wallKick(g.Board, r) {
			g.CurrBlock.rotateClockWise()
		}
	case Drop:
		drop := true
		for drop {
			g.CurrBlock.y++
			drop = !g.Board.collide(g.CurrBlock)
		}
		g.CurrBlock.y--
		g.nextStep()
	case Escape:
		g.State = Over
	case Pause:
		g.State = Paused
	}
}

func (g *Game) handleKeyPaused(msg Msg) {
	switch msg.Key {
	case Pause:
		g.State = Playing
	}
}

func (g *Game) handleKeyGameOver(msg Msg) {
}

func (g *Game) updatePlaying(delta int64) {
	g.stepTimer += delta
	if g.stepTimer > g.speed() {
		if g.step(g.Board, g.CurrBlock) {
			_ = g.Board.removeLines()
			//todo : score

			g.setNowToNext()

			if g.Board.collide(g.CurrBlock) {
				g.State = Over
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
