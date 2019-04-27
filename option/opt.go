package option

// Opt 구조체는 게임 실행과 관련 된 옵션을 담는다.
// X = width, Y = height
// CellSize = 격자 하나의 크기
// Title = 윈도우 이름
type Opt struct {
	X, Y     int
	CellSize int
	Title    string
}
