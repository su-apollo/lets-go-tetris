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

// Cell 타입은 Board 자료구조의 2차원 좌표가 비어있는지 여부를 관리한다.
type Cell bool

// Color 타입은 색상 정보를 3원색 기반으로 표현한다.
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

// Block 인터페이스는 Board 위에서 동작하는 Tetromino 블록의 동작을 서술한다.
type Block interface {
	GetCells() [][]Cell
	GetPosition() (int, int)
	GetColor() Color
}

// Board 인터페이스는 Cell 자료구조를 2차원 좌표로 배열하고, 해당 좌표의 색상 값을 반환하는 동작을 서술한다.
type Board interface {
	GetCells() [][]Cell
	GetColor(x int, y int) Color
}
