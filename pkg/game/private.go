package game

// Rotation 타입은 미노의 회전 가능한 유형을 나타낸다.
type Rotation int

// Rotate 타입은 미노의 회전했을 때 경우의 수를 나타낸다.
type Rotate int

// Rotation types
const (
	ZeroRotation Rotation = 0 + iota
	RightRotation
	TwoRotation
	LeftRotation
	RotationMax
)

// Rotate types
const (
	ZtoR Rotate = 0 + iota
	RtoZ
	RtoT
	TtoR
	TtoL
	LtoT
	LtoZ
	ZtoL
)

const (
	shapeX = 4
	shapeY = 4
)

var shapes = [][]string{
	{i0, i1, i2, i3},
	{j0, j1, j2, j3},
	{l0, l1, l2, l3},
	{o0, o0, o0, o0},
	{s0, s1, s2, s3},
	{t0, t1, t2, t3},
	{z0, z1, z2, z3},
}

var wallKicks = map[Rotate][][]int{
	ZtoR: {{0, 0}, {-1, 0}, {-1, -1}, {0, +2}, {-1, +2}},
	RtoZ: {{0, 0}, {+1, 0}, {+1, +1}, {0, -2}, {+1, -2}},
	RtoT: {{0, 0}, {+1, 0}, {+1, +1}, {0, -2}, {+1, -2}},
	TtoR: {{0, 0}, {-1, 0}, {-1, -1}, {0, +2}, {-1, +2}},
	TtoL: {{0, 0}, {+1, 0}, {+1, -1}, {0, +2}, {+1, +2}},
	LtoT: {{0, 0}, {-1, 0}, {-1, +1}, {0, -2}, {-1, -2}},
	LtoZ: {{0, 0}, {-1, 0}, {-1, +1}, {0, -2}, {-1, -2}},
	ZtoL: {{0, 0}, {+1, 0}, {+1, -1}, {0, +2}, {+1, +2}},
}

var iKicks = map[Rotate][][]int{
	ZtoR: {{0, 0}, {-2, 0}, {+1, 0}, {-2, +1}, {+1, -2}},
	RtoZ: {{0, 0}, {+2, 0}, {-1, 0}, {+2, -1}, {-1, +2}},
	RtoT: {{0, 0}, {-1, 0}, {+2, 0}, {-1, -2}, {+2, +1}},
	TtoR: {{0, 0}, {+1, 0}, {-2, 0}, {+1, +2}, {-2, -1}},
	TtoL: {{0, 0}, {+2, 0}, {-1, 0}, {+2, -1}, {-1, +2}},
	LtoT: {{0, 0}, {-2, 0}, {+1, 0}, {-2, +1}, {+1, -2}},
	LtoZ: {{0, 0}, {+1, 0}, {-2, 0}, {+1, +2}, {-2, -1}},
	ZtoL: {{0, 0}, {-1, 0}, {+2, 0}, {-1, -2}, {+2, +1}},
}

const i0 = `
oooo
xxxx
oooo
oooo
`

const i1 = `
ooxo
ooxo
ooxo
ooxo
`

const i2 = `
oooo
oooo
xxxx
oooo
`

const i3 = `
oxoo
oxoo
oxoo
oxoo
`

const j0 = `
xooo
xxxo
oooo
oooo
`

const j1 = `
oxxo
oxoo
oxoo
oooo
`

const j2 = `
oooo
xxxo
ooxo
oooo
`

const j3 = `
oxoo
oxoo
xxoo
oooo
`

const s0 = `
oxxo
xxoo
oooo
oooo
`

const s1 = `
oxoo
oxxo
ooxo
oooo
`

const s2 = `
oooo
oxxo
xxoo
oooo
`

const s3 = `
xooo
xxoo
oxoo
oooo
`

const z0 = `
xxoo
oxxo
oooo
oooo
`

const z1 = `
ooxo
oxxo
oxoo
oooo
`

const z2 = `
oooo
xxoo
oxxo
oooo
`
const z3 = `
oxoo
xxoo
xooo
oooo
`

const t0 = `
oxoo
xxxo
oooo
oooo
`

const t1 = `
oxoo
oxxo
oxoo
oooo
`
const t2 = `
oooo
xxxo
oxoo
oooo
`

const t3 = `
oxoo
xxoo
oxoo
oooo
`

const o0 = `
oxxo
oxxo
oooo
oooo
`

const l0 = `
ooxo
xxxo
oooo
oooo
`

const l1 = `
oxoo
oxoo
oxxo
oooo
`

const l2 = `
oooo
xxxo
xooo
oooo
`

const l3 = `
xxoo
oxoo
oxoo
oooo
`
