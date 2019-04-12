package sdl2

import "github.com/veandco/go-sdl2/sdl"

type Wrapper struct {
}

func (*Wrapper) Init() error {
	return sdl.Init(sdl.INIT_EVERYTHING)
}

func (*Wrapper) Quit() {
	sdl.Quit()
}
