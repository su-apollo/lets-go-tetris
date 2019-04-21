package mock

import (
	"lets-go-tetris/interfaces/renderer"
)

// Event 타입은 renderer 패키지의 Event 를 mocking 한다.
type Event renderer.Event

// Render 타입은 renderer 패키지의 Render 를 mocking 한다.
type Render renderer.Render

// Window 타입은 renderer 패키지의 Window 를 mocking 한다.
type Window renderer.Window

// Surface 타입은 renderer 패키지의 Surface 를 mocking 한다.
type Surface renderer.Surface
