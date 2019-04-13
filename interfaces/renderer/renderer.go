package renderer

import (
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/event"
)

type Info interface {
	GetPos() (int32, int32)
	GetColor() uint32
}

type Renderer interface {
	Render([]Info) error
	Update() ([]event.Msg, bool)
}

type Object interface {
	RenderInfo() []Info
}

type Event interface {
	GetType() uint32      // GetType returns the event type
	GetTimestamp() uint32 // GetTimestamp returns the timestamp of the event
}

type Render interface {
	Init() error
	Quit()
	PollEvent() Event
	Update()
}

type Window interface {
	GetSurface()
	UpdateSurface() *sdl.Surface
	Destroy()
}

type Surface interface {
	FillRect(rect *sdl.Rect, color uint32) error
}
