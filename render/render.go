package render

import (
	"github.com/veandco/go-sdl2/sdl"
	_ "github.com/veandco/go-sdl2/sdl"
	"lets-go-tetris/option"
)

type Object interface {
	Render() bool
	Next() bool
}

func NewWrapper(opt option.Opt) (*SDLWrapper, error) {
	wrapper := &SDLWrapper{opt, nil, nil, nil}

	err := wrapper.init()
	return wrapper, err
}

type fn func()

type SDLWrapper struct {
	opt option.Opt

	deferFn   []fn
	destroyFn []fn

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

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	wrapper.surface = surface

	wrapper.destroyFn = wrapper.deferFn
	wrapper.deferFn = wrapper.deferFn[:]
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
