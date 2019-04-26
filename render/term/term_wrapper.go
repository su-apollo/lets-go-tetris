package term

import (
	"github.com/nsf/termbox-go"
	"lets-go-tetris/event"
	"lets-go-tetris/option"
	"lets-go-tetris/render"
)

func NewTermWrapper(opt option.Opt) (*Wrapper, error) {
	wrapper := &Wrapper{opt: opt}

	err := wrapper.init()
	return wrapper, err
}

type Wrapper struct {
	opt option.Opt

	width, height int
	backBuffer []termbox.Cell
}

func (wrapper *Wrapper) init() error {
	err := termbox.Init()
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	wrapper.resetBackBuffer(termbox.Size())

	if err != nil {
		return err
	}

	return nil
}

func (wrapper *Wrapper) resetBackBuffer(w, h int) {
	wrapper.width = w
	wrapper.height = h
	wrapper.backBuffer = make([]termbox.Cell, w*h)
}

func (wrapper *Wrapper) clear() error {
	return termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func (wrapper *Wrapper) flush() error {
	return termbox.Flush()
}

func (wrapper *Wrapper) Render(info []render.Info) error {
	if err := wrapper.clear(); err != nil {
		return err
	}

	for _, i := range info {
		posX, posY := i.GetPos()
		termbox.SetCell(int(posX), int(posY), ' ', termbox.ColorBlack, termbox.ColorWhite)
	}

	if err := wrapper.flush(); err != nil {
		return err
	}

	return nil
}

func (wrapper *Wrapper) Update() ([]event.Msg, bool) {
	var keys []event.Msg

	switch e := termbox.PollEvent(); e.Type {
	case termbox.EventKey:
		if msg, ok := termKeyCodeToEvent(e.Key); ok {
			keys = append(keys, msg)
		}
	}

	return keys, true
}

func termKeyCodeToEvent(k termbox.Key) (event.Msg, bool) {
	switch k {
	case termbox.KeyArrowLeft:
		return event.Msg{Key: event.Left}, true
	case termbox.KeyArrowRight:
		return event.Msg{Key: event.Right}, true
	case termbox.KeyArrowUp:
		return event.Msg{Key: event.Up}, true
	case termbox.KeyArrowDown:
		return event.Msg{Key: event.Down}, true
	case termbox.KeyEsc:
		return event.Msg{Key: event.Escape}, true
	default:
		return event.Msg{Key: event.Nop}, false
	}
}