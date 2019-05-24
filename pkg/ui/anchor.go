package ui

// type Stretch int
type NineGrid int

const (
	LeftTop NineGrid = iota
	Top
	RightTop
	Left
	Center
	Right
	LeftBottom
	Bottom
	RightBottom
)

type Anchor interface {
	Position(uint, uint) (int, int)
}

func NewPresetAnchor(g NineGrid) Anchor {
	return &presetAnchor{g}
}

type presetAnchor struct {
	grid NineGrid
}

// Calculate position by screen width, height
func (a *presetAnchor) Position(w uint, h uint) (int, int) {
	return 0, 0
}
