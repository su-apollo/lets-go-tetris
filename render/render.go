package render

import (
	"lets-go-tetris/event"
)

type Info struct {
	PosX, PosY int32

	Color uint32
}

type Renderer interface {
	Render(Object)
	Update() ([]event.Msg, bool)
}

type Object interface {
	RenderInfo() []Info
}
