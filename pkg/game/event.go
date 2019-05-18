package game

// Key type is the type of event that can occur as input in a game.
type Key int

// Key types
const (
	Left Key = iota
	Right
	Down
	Drop
	ClockWise
	CounterClockWise
	Escape
	Pause
	Nop
)

// Msg 타입은 게임에서 발생할 수 있는 이벤트를 담는 자료구조다.
type Msg struct {
	Key
}
