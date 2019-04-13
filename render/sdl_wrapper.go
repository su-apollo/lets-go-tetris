package render

import (
	"github.com/veandco/go-sdl2/sdl"

	"lets-go-tetris/event"
	"lets-go-tetris/option"
)

func NewSDLWrapper(opt option.Opt) (*SDLWrapper, error) {
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

func (wrapper *SDLWrapper) Clear() error {
	return wrapper.surface.FillRect(nil, 0x000000)
}

func (wrapper *SDLWrapper) Render(info []Info) error {
	if err := wrapper.Clear(); err != nil {
		return err
	}

	for _, i := range info {
		r := sdl.Rect{
			X: i.PosX * wrapper.opt.CellSize,
			Y: i.PosY * wrapper.opt.CellSize,
			W: wrapper.opt.CellSize,
			H: wrapper.opt.CellSize,
		}
		_ = wrapper.surface.FillRect(&r, i.Color)
	}
	return nil
}

func (wrapper *SDLWrapper) Update() ([]event.Msg, bool) {
	wrapper.window.UpdateSurface()

	var keys []event.Msg
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch t := e.(type) {
		case *sdl.QuitEvent:
			return nil, false
		case *sdl.KeyboardEvent:
			if e.GetType() == sdl.KEYDOWN {
				if msg, ok := sdlKeyCodeToEvent(t.Keysym.Sym); ok {
					keys = append(keys, msg)
				}
			}
		}
	}
	return keys, true
}

func sdlKeyCodeToEvent(k sdl.Keycode) (event.Msg, bool) {
	switch k {
	case sdl.K_LEFT, sdl.K_a, sdl.K_j:
		return event.Msg{Key: event.Left}, true
	case sdl.K_RIGHT, sdl.K_d, sdl.K_l:
		return event.Msg{Key: event.Right}, true
	case sdl.K_UP, sdl.K_w, sdl.K_i:
		return event.Msg{Key: event.Up}, true
	case sdl.K_DOWN, sdl.K_s, sdl.K_k, sdl.K_SPACE:
		return event.Msg{Key: event.Down}, true
	case sdl.K_ESCAPE:
		return event.Msg{Key: event.Escape}, true
	case sdl.K_p:
		return event.Msg{Key: event.Pause}, true
	default:
		return event.Msg{Key: event.Nop}, false
	}
}
