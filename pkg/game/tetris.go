package game

import (
	"math/rand"
	"time"
)

const startX = 3

// Yes, This game is tetirs
type tetris struct {
	state State
	now   *tetromino
	next  *tetromino
	keep  *tetromino
	stack []Shape
	mat   *matrix

	stepTimer int64
}

//
func (t tetris) State() State {
	return t.state
}

//
func (t tetris) NowBlock() Block {
	return t.now
}

//
func (t tetris) NextBlock() Block {
	return t.next
}

//
func (t tetris) KeepBlock() Block {
	return t.keep
}

//
func (t tetris) GhostBlock() Block {
	ghost := &ghost{}
	ghost.init(t.mat, t.now)
	return ghost
}

//
func (t tetris) Board() Board {
	return t.mat
}

//
func (t *tetris) HandleKey(msg Msg) {
	switch t.state {
	case Playing:
		t.handleKeyPlaying(msg)
		break
	case Paused:
		t.handleKeyPaused(msg)
		break
	case Over:
		t.handleKeyGameOver(msg)
		break
	}
}

func (t *tetris) Update(delta int64) {
	switch t.state {
	case Playing:
		t.updatePlaying(delta)
		break
	case Paused:
		t.updatePaused(delta)
		break
	case Over:
		t.updateGameOver(delta)
		break
	}
}

func (t *tetris) setNowToNext() {
	t.now = t.next
	t.now.x = startX

	s := t.popQueue()
	t.next = newTetromino(s)

	if t.stack == nil {
		t.resetStack()
		t.shuffleStack()
	}
}

func (t *tetris) reset() {
	t.mat.reset()
	t.resetStack()
	t.shuffleStack()

	s := t.popQueue()
	t.now = newTetromino(s)
	t.now.x = startX

	s = t.popQueue()
	t.next = newTetromino(s)

	t.state = Playing
}

func (t *tetris) resetStack() {
	t.stack = []Shape{I, J, L, O, S, T, Z}
}

func (t *tetris) shuffleStack() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(t.stack); n > 0; n-- {
		i := r.Intn(n)
		t.stack[n-1], t.stack[i] = t.stack[i], t.stack[n-1]
	}
}

func (t *tetris) popQueue() Shape {
	n := len(t.stack)
	if n < 2 {
		v := t.stack[0]
		t.stack = nil
		return v
	} else {
		v := t.stack[n-1]
		t.stack = t.stack[:n-1]
		return v
	}
}

func (t *tetris) handleKeyPlaying(msg Msg) {
	switch msg.Key {
	case Left:
		t.now.x--
		if t.mat.Collide(t.now) {
			t.now.x++
		}
	case Right:
		t.now.x++
		if t.mat.Collide(t.now) {
			t.now.x--
		}
	case Down:
		t.now.y++
		if t.mat.Collide(t.now) {
			t.now.y--
			t.nextStep()
		}
	case ClockWise:
		r := t.now.rotateClockWise()
		if !t.now.wallKick(t.mat, r) {
			t.now.rotateCounterClockWise()
		}
	case CounterClockWise:
		r := t.now.rotateCounterClockWise()
		if !t.now.wallKick(t.mat, r) {
			t.now.rotateClockWise()
		}
	case Drop:
		drop := true
		for drop {
			t.now.y++
			drop = !t.mat.Collide(t.now)
		}
		t.now.y--
		t.nextStep()
	case Escape:
		t.state = Over
	case Pause:
		t.state = Paused
	}
}

func (t *tetris) handleKeyPaused(msg Msg) {
	switch msg.Key {
	case Pause:
		t.state = Playing
	}
}

func (t *tetris) handleKeyGameOver(msg Msg) {
}

func (t *tetris) updatePlaying(delta int64) {
	t.stepTimer += delta
	if t.stepTimer > t.speed() {
		if t.step() {
			_ = t.mat.removeLines()
			//todo : score

			t.setNowToNext()

			if t.mat.Collide(t.now) {
				t.state = Over
			}
		}
		t.stepTimer = 0
	}
}

func (t *tetris) updatePaused(delta int64) {
}

func (t *tetris) updateGameOver(delta int64) {
	t.reset()
}

func (t *tetris) speed() int64 {
	// todo : game level
	return 1000000000
}

func (t *tetris) nextStep() {
	t.stepTimer += t.speed()
}

func (t *tetris) step() bool {
	t.now.y++
	if !t.mat.Collide(t.now) {
		return false
	}

	t.now.y--
	t.mat.merge(t.now)

	return true
}
