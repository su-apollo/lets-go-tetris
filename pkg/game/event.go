package game

// Key 타입은 게임에서 입력으로 발생할 수 있는 이벤트의 종류다.
type Key int

// Left  			오른쪽 이동
// Right 			왼쪽 이동
// Down 			아래로 이동
// Drop				빠르게 밑으로 이동
// ClockWise		시계방향 회전
// CounterClockWise	반시계방향 회전
// Escape 			일시 중지
// Pause 			없음
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
