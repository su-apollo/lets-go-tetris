package event

type Key int

const (
	Left Key = iota
	Right
	Up
	Down
	Escape
	Pause
	Nop
)

type Msg struct {
	Key
}
