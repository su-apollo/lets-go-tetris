package render

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"lets-go-tetris/mock/mockfakes"
	"lets-go-tetris/option"
)

var _ = XDescribe("SDL2 Wrapper 테스트", func() {
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

var _ = Describe("SDL2 Wrapper mock 의존성 주입", func() {
	It("아직 얼개 짜기 중...", func() {
		var init bool
		wrapper := mockfakes.FakeRender{}
		wrapper.InitCalls(func() error {
			init = true
			return nil
		})

		wrapper.QuitCalls(func() {
			init = false
		})

		Expect(wrapper.Init()).ShouldNot(HaveOccurred())
		Expect(init).Should(BeTrue())

		wrapper.QuitStub()
		Expect(init).Should(BeFalse())
	})
})
