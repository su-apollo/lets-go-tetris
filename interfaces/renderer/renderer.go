package renderer

import (
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/event"
)

// Info 인터페이스는 렌더링에 필요한 정보(info)를 반환한다.
type Info interface {
	GetPos() (int32, int32)
	GetColor() uint32
}

// Renderer 인터페이스는 화면에 렌더링을 담당한다.
type Renderer interface {
	Render([]Info) error
	Update() ([]event.Msg, bool)
}

// Object 인터페이스는 렌더링 정보를 반환하는 RenderInfo() 함수를 구현한다.
type Object interface {
	RenderInfo() []Info
}

// Event 인터페이스는 SDL2 의 Event 타입을 mocking 한다.
type Event interface {
	GetType() uint32      // GetType returns the event type
	GetTimestamp() uint32 // GetTimestamp returns the timestamp of the event
}

// Render 인터페이스는 SDL2 의 함수 전반을 mocking 한다.
type Render interface {
	Init() error
	Quit()
	CreateWindow(string, int32, int32, int32, int32, uint32) (*Window, error)
	PollEvent() Event
	Update()
}

// Window 인터페이스는 SDL2가 wrapping 하고 있는 Window 타입을 mocking 한다.
type Window interface {
	GetSurface() (*sdl.Surface, error)
	UpdateSurface() error
	Destroy() error
}

// Surface 인터페이스는 SDL2가 wrapping 하고 있는 Surface 타입을 mocking 한다.
type Surface interface {
	FillRect(rect *sdl.Rect, color uint32) error
}
