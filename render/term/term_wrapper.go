package term

import (
	"github.com/nsf/termbox-go"
	"lets-go-tetris/option"
)

func NewTermWrapper(opt option.Opt) (*Wrapper, error) {
	wrapper := &Wrapper{opt: opt}

	err := wrapper.init()
	return wrapper, err
}

type Wrapper struct {
	opt option.Opt

}

func (wrapper *Wrapper) init() error {
	err := termbox.Init()
	defer termbox.Close()

	if err != nil {
		return err
	}

	return nil
}