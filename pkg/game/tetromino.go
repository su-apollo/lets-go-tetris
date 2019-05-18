package game

type tetromino struct {
	shape    Shape
	x, y     int
	cells    [][][]Cell
	rotation Rotation
}

func (t tetromino) Cells() [][]Cell {
	return t.cells[t.rotation]
}

func (t tetromino) Position() (int, int) {
	return t.x, t.y
}

func (t tetromino) Shape() Shape {
	return t.shape
}

func newTetromino(s Shape) *tetromino {
	t := &tetromino{}
	t.init(shapes[s])
	t.shape = s
	return t
}

func (t *tetromino) init(rotationShapes []string) {
	t.cells = make([][][]Cell, RotationMax)
	for i := range t.cells {
		t.cells[i] = make([][]Cell, shapeY)
		for j := range t.cells[i] {
			t.cells[i][j] = make([]Cell, shapeX)
		}
	}

	for r, shape := range rotationShapes {
		i := 0
		j := 0
		for _, c := range shape {
			switch c {
			case 'x':
				t.cells[r][i][j] = true
				fallthrough
			case 'o':
				j++
				if j >= shapeY {
					i++
					j = 0
				}
			}
		}
	}
}

func (t *tetromino) rotateClockWise() Rotate {
	var r Rotate
	switch t.rotation {
	case ZeroRotation:
		r = ZtoR
	case RightRotation:
		r = RtoT
	case TwoRotation:
		r = TtoL
	case LeftRotation:
		r = LtoZ
	}
	t.rotate(t.rotation + 1)
	return r
}

func (t *tetromino) rotateCounterClockWise() Rotate {
	var r Rotate
	switch t.rotation {
	case ZeroRotation:
		r = ZtoL
	case RightRotation:
		r = RtoZ
	case TwoRotation:
		r = TtoR
	case LeftRotation:
		r = LtoT
	}
	t.rotate(t.rotation - 1)
	return r
}

func (t *tetromino) rotate(r Rotation) {
	t.rotation = (r%RotationMax + RotationMax) % RotationMax
}

func (t *tetromino) wallKick(b Board, r Rotate) bool {
	if t.shape == I {
		for _, v := range iKicks[r] {
			t.x += v[0]
			t.y += v[1]
			if !b.Collide(t) {
				return true
			}
			t.x -= v[0]
			t.y -= v[1]
		}
	} else {
		for _, v := range wallKicks[r] {
			t.x += v[0]
			t.y += v[1]
			if !b.Collide(t) {
				return true
			}
			t.x -= v[0]
			t.y -= v[1]
		}
	}
	return false
}
