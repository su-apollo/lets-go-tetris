package mocks

//go:generate counterfeiter -o event.go --fake-name Event ../render Event
//go:generate counterfeiter -o render.go --fake-name Render ../render Render
//go:generate counterfeiter -o window.go --fake-name Window ../render Window
//go:generate counterfeiter -o surface.go --fake-name Surface ../render Surface
