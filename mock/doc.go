package mock

//go:generate gobin -m -run github.com/maxbrunsfeld/counterfeiter/v6 . Event
//go:generate gobin -m -run github.com/maxbrunsfeld/counterfeiter/v6 . Render
//go:generate gobin -m -run github.com/maxbrunsfeld/counterfeiter/v6 . Window
//go:generate gobin -m -run github.com/maxbrunsfeld/counterfeiter/v6 . Surface
