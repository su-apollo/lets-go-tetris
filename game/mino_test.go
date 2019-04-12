package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"lets-go-tetris/render"
	"time"
)

var _ = XDescribe("mino 초기화 성공 테스트", func() {
	type testData struct {
		input    string
		expected []cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		m := mino{}
		m.init(d.input)
		actual := m.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("", testData{S, []cell{
			false, true, true,
			true, true, false,
			false, false, false,
		}}),
		Entry("", testData{Z, []cell{
			true, true, false,
			false, true, true,
			false, false, false,
		}}),
		Entry("", testData{T, []cell{
			false, true, false,
			true, true, true,
			false, false, false,
		}}),
		Entry("", testData{I, []cell{
			false, true, false, false,
			false, true, false, false,
			false, true, false, false,
			false, true, false, false,
		}}),
		Entry("", testData{O, []cell{
			true, true,
			true, true,
		}}),
		Entry("", testData{L, []cell{
			false, false, true,
			true, true, true,
			false, false, false,
		}}),
		Entry("", testData{J, []cell{
			true, false, false,
			true, true, true,
			false, false, false,
		}}),
	)
})

var _ = XDescribe("mino srs 테스트 (super rotate system)", func() {
	type testData struct {
		inputInit    	string
		inputRotate		int
		expected 		[]cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		m := mino{}
		m.init(d.inputInit)
		m.rotate(d.inputRotate)
		actual := m.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("", testData{S, -1, []cell{
			true, false, false,
			true, true, false,
			false, true, false,
		}}),
		Entry("", testData{Z, 6, []cell{
			false, false, false,
			true, true, false,
			false, true, true,
		}}),
		Entry("", testData{T, -2,[]cell{
			false, false, false,
			true, true, true,
			false, true, false,
		}}),
		Entry("", testData{I, 16,[]cell{
			false, false, false, false,
			true, true, true, true,
			false, false, false, false,
			false, false, false, false,
		}}),
		Entry("", testData{O, -4,[]cell{
			true, true,
			true, true,
		}}),
		Entry("", testData{L, 7,[]cell{
			true, true, false,
			false, true, false,
			false, true, false,
		}}),
		Entry("", testData{J, -17,[]cell{
			false, true, false,
			false, true, false,
			true, true, false,
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
