package event

// Key 타입은 게임에서 입력으로 발생할 수 있는 이벤트의 종류다.
type Key int

// Left  	오른쪽 이동
// Right 	회전
// Up 		빠르게 내리기
// Down 	메뉴
// Escape 	일시 중지
// Pause 	없음
const (
	Left Key = iota
	Right
	Up
	Down
	Escape
	Pause
	Nop
)

// Msg 타입은 게임에서 발생할 수 있는 이벤트를 담는 자료구조다.
type Msg struct {
	Key
}
