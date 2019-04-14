package render

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"lets-go-tetris/interfaces/renderer"
	"lets-go-tetris/mock/mockfakes"
)

var _ = Describe("SDL2 Wrapper 테스트", func() {
	var wrapper *SDLWrapper
	var window *mockfakes.FakeWindow
	surface := &mockfakes.FakeSurface{}

	BeforeEach(func() {
		window = &mockfakes.FakeWindow{}
		surface = &mockfakes.FakeSurface{}

		wrapper = &SDLWrapper{}
		wrapper.window = window
		wrapper.surface = surface
	})

	AfterEach(func() {
		wrapper.Close()
		wrapper = nil
	})

	It("Render 함수가 그려야 할 정보를 잘 전달한다.", func() {
		Expect(surface.FillRectCallCount()).Should(Equal(0))
		err := wrapper.Render([]renderer.Info{})
		Expect(err).ShouldNot(HaveOccurred())

		// +1 : clear call 1 + draw call 0 = 1
		Expect(surface.FillRectCallCount()).Should(Equal(1))

		err = wrapper.Render([]renderer.Info{
			&InfoImpl{12, 34, 5678},
		})
		Expect(err).ShouldNot(HaveOccurred())

		// +2 : clear call 1 + draw call 1 = 2
		Expect(surface.FillRectCallCount()).Should(Equal(3))
	})
})
