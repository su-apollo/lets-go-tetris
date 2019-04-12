package render

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"lets-go-tetris/option"
)

var _ = Describe("SDL2 Wrapper 테스트", func() {
	var wrapper *SDLWrapper
	opt := option.Opt{X: 123, Y: 123, CellSize: 123, Title: "test"}

	BeforeEach(func() {
		var err error
		wrapper, err = NewSDLWrapper(opt)
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		wrapper.Close()
		wrapper = nil
	})
})
