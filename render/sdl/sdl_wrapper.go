package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/event"
	"lets-go-tetris/option"
	"lets-go-tetris/render"
)

// NewSDLWrapper 함수는 Renderer 인터페이스를 구현한, SDL2 Wrapper 구조체를 반환한다.
func NewSDLWrapper(opt option.Opt) (*Wrapper, error) {
	wrapper := &Wrapper{opt: opt}

	err := wrapper.init()
	return wrapper, err
}

type fn func()

// Wrapper 구조체는 SDL2 라이브러리를 감싸고, Renderer 인터페이스를 구현한다.
type Wrapper struct {
	opt option.Opt

	deferFn   []fn
	destroyFn []fn

	window   render.Window
	//renderer render.Renderer
	surface  render.Surface
}

const shapeX = 4

func (wrapper *Wrapper) init() error {
	defer func() {
		for _, f := range wrapper.deferFn {
			f()
		}
	}()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	wrapper.pushFn(sdl.Quit)

	width := (wrapper.opt.X + shapeX) * wrapper.opt.CellSize
	height := wrapper.opt.Y * wrapper.opt.CellSize
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

	//renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	//wrapper.renderer = renderer

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	surface.SetBlendMode(sdl.BLENDMODE_BLEND)
	wrapper.surface = surface

	wrapper.destroyFn = wrapper.deferFn
	wrapper.deferFn = nil
	return nil
}

func (wrapper *Wrapper) pushFn(f fn) {
	wrapper.deferFn = append([]fn{func() {
		sdl.Quit()
	}}, wrapper.deferFn...)
}

// Close SDLWrapper 내부에서 할당한 자원이 있다면 적절히 해제한다.
func (wrapper *Wrapper) Close() {
	for _, f := range wrapper.destroyFn {
		f()
	}
}

func (wrapper *Wrapper) clear() error {
	return wrapper.surface.FillRect(nil, 0x000000)
}

// Render 함수는 인자로 전달받은 렌더링 관련 정보를 적절히 해독하여 화면에 출력한다.
func (wrapper *Wrapper) Render(info []render.Info) error {
	if err := wrapper.clear(); err != nil {
		return err
	}

	for _, i := range info {
		posX, posY := i.GetPos()
		r := sdl.Rect{
			X: int32(posX * wrapper.opt.CellSize),
			Y: int32(posY * wrapper.opt.CellSize),
			W: int32(wrapper.opt.CellSize),
			H: int32(wrapper.opt.CellSize),
		}
		_ = wrapper.surface.FillRect(&r, i.GetColor())
	}
	return nil
}

// Update 함수는 화면을 적절히 갱신 한 후, 키보드 입력을 처리한다.
func (wrapper *Wrapper) Update() ([]event.Msg, bool) {
	if err := wrapper.window.UpdateSurface(); err != nil {
		return nil, false
	}

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
		return event.Msg{Key: event.ClockWise}, true
	case sdl.K_DOWN, sdl.K_s, sdl.K_k:
		return event.Msg{Key: event.Down}, true
	case sdl.K_SPACE:
		return event.Msg{Key: event.Drop}, true
	case sdl.K_ESCAPE:
		return event.Msg{Key: event.Escape}, true
	case sdl.K_p:
		return event.Msg{Key: event.Pause}, true
	default:
		return event.Msg{Key: event.Nop}, false
	}
}
