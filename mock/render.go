package mock

type Render interface {
	Init() error
	Quit()
}
