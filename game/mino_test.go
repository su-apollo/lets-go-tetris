package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"lets-go-tetris/render"
	"time"
)

var _ = Describe("mino 초기화 성공 테스트", func() {
	type testData struct {
		input    string
		expected []cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		b := mino{}
		b.init(d.input)
		actual := b.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("", testData{S, []cell{
			false, true, true, false,
			true, true, false, false,
			false, false, false, false,
			false, false, false, false,
		}}),
		Entry("", testData{Z, []cell{
			true, true, false, false,
			false, true, true, false,
			false, false, false, false,
			false, false, false, false,
		}}),
		Entry("", testData{T, []cell{
			false, true, false, false,
			true, true, true, false,
			false, false, false, false,
			false, false, false, false,
		}}),
		Entry("", testData{I, []cell{
			false, true, false, false,
			false, true, false, false,
			false, true, false, false,
			false, true, false, false,
		}}),
	)
})

var _ = Describe("random 통제 테스트", func() {
	It("seed 값이 같으면 동일한 결과가 나온다.", func() {
		expected := NewRandomMino(0)
		actual := NewRandomMino(0)
		diff := deep.Equal(expected.cells, actual.cells)
		Expect(diff).Should(BeNil())
	})

	It("seed 값이 다르면 결과도 다르게.", func() {
		expected := NewRandomMino(time.Now().UnixNano())
		actual := NewRandomMino(time.Now().UnixNano() + 1)
		diff := deep.Equal(expected.cells, actual.cells)
		Expect(diff).ShouldNot(BeNil())
	})
})

var _ = Describe("mino.RenderInfo() 함수가", func() {
	It("렌더링 정보를 제대로 반환한다.", func() {
		m := mino{
			cells:  []cell{false, true, false, true},
			color:  123,
			x: 		1234,
			y:		4321,
		}
		expected := []render.Info{
			{PosX: 1 + 1234, PosY: 4321, Color: 123},
			{PosX: 3 + 1234, PosY: 4321, Color: 123},
		}
		actual := m.RenderInfo()
		Expect(actual).Should(Equal(expected))
	})
})
