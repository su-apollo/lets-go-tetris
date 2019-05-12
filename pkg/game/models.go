package game

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

type Cell bool

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

const (
	o Cell = true
	x Cell = false
)

type Block interface {
	GetCells() [][]Cell
	GetPosition() (int, int)
	GetColor() Color
}

type Board interface {
	GetCells() [][]Cell
	GetColor(x int, y int) Color
}
