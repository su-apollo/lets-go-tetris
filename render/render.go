package render

import (
	"github.com/veandco/go-sdl2/sdl"
	_ "github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/option"
)

type Info struct {
	PosX, PosY int32

	Color uint32
}

type Renderer interface {
	Render(Object)
	Update() ([]sdl.Keycode, bool)
}

type Object interface {
	RenderInfo() []Info
}

func NewWrapper(opt option.Opt) (*SDLWrapper, error) {
	wrapper := &SDLWrapper{opt: opt}

	err := wrapper.init()
	return wrapper, err
}

type fn func()

type SDLWrapper struct {
	opt option.Opt

	deferFn   []fn
	destroyFn []fn

	window  *sdl.Window
	surface *sdl.Surface
}

const shapeX = 4

func (wrapper *SDLWrapper) init() error {
	defer func() {
		for _, f := range wrapper.deferFn {
			f()
		}
	}()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	wrapper.pushFn(sdl.Quit)

	width := int32(wrapper.opt.X+shapeX) * wrapper.opt.CellSize
	height := int32(wrapper.opt.Y) * wrapper.opt.CellSize
	window, err := sdl.CreateWindow(
		"lets go",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(width),
		int32(height),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return err
	}
	wrapper.pushFn(func() { window.Destroy() })
	wrapper.window = window

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	wrapper.surface = surface

	wrapper.destroyFn = wrapper.deferFn
	wrapper.deferFn = nil
	return nil
}

func (wrapper *SDLWrapper) pushFn(f fn) {
	wrapper.deferFn = append([]fn{func() {
		sdl.Quit()
	}}, wrapper.deferFn...)
}

func (wrapper *SDLWrapper) Close() {
	for _, f := range wrapper.destroyFn {
		f()
	}
}

func (wrapper *SDLWrapper) Render(o Object) {
	for _, info := range o.RenderInfo() {
		r := sdl.Rect{
			X: info.PosX * wrapper.opt.CellSize,
			Y: info.PosY * wrapper.opt.CellSize,
			W: wrapper.opt.CellSize,
			H: wrapper.opt.CellSize,
		}
		_ = wrapper.surface.FillRect(&r, info.Color)
	}
}

func (wrapper *SDLWrapper) Update() ([]sdl.Keycode, bool) {
	wrapper.window.UpdateSurface()

	var keys []sdl.Keycode
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return nil, false
		case *sdl.KeyboardEvent:
			if event.GetType() == sdl.KEYDOWN {
				keys = append(keys, t.Keysym.Sym)
			}
		}
	}
	return keys, true
}
