package mock

type Event interface {
	GetType() uint32      // GetType returns the event type
	GetTimestamp() uint32 // GetTimestamp returns the timestamp of the event
}

type Render interface {
	Init() error
	Quit()
	PollEvent() Event
}
