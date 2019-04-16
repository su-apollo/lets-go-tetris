package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"lets-go-tetris/interfaces/renderer"
	"lets-go-tetris/render"
	"math/rand"
)

var _ = Describe("mino 초기화 성공 테스트", func() {
	type testData struct {
		input    []string
		expected []cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		m := mino{}
		m.init(d.input)
		actual := m.currentCells()
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("S", testData{shapes[S], []cell{
			false,	true,	true,	false,
			true,	true,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("Z", testData{shapes[Z], []cell{
			true,	true,	false,	false,
			false,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("T", testData{shapes[T], []cell{
			false,	true,	false,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("I", testData{shapes[I], []cell{
			false,	false,	false,	false,
			true,	true,	true,	true,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("O", testData{shapes[O], []cell{
			false,	true,	true,	false,
			false,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("L", testData{shapes[L], []cell{
			false,	false,	true,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("J", testData{shapes[J], []cell{
			true,	false,	false,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
	)
})

var _ = Describe("mino srs 테스트 (super rotate system)", func() {
	type testData struct {
		inputInit		[]string
		inputRotation	int
		expected		[]cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		m := mino{}
		m.init(d.inputInit)
		m.rotate(d.inputRotation)
		actual := m.currentCells()
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("S", testData{shapes[S], -1, []cell{
			true,	false,	false,	false,
			true,	true,	false,	false,
			false,	true,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("Z", testData{shapes[Z], 6, []cell{
			false,	false,	false,	false,
			true,	true,	false,	false,
			false,	true,	true,	false,
			false,	false,	false,	false,
		}}),
		Entry("T", testData{shapes[T], -2, []cell{
			false,	false,	false,	false,
			true,	true,	true,	false,
			false,	true,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("I", testData{shapes[I], 16, []cell{
			false,	false,	false,	false,
			true,	true,	true,	true,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("O", testData{shapes[O], -4, []cell{
			false,	true,	true,	false,
			false,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("L", testData{shapes[L], 7, []cell{
			true,	true,	false,	false,
			false,	true,	false,	false,
			false,	true,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("J", testData{shapes[J], -17, []cell{
			false,	true,	false,	false,
			false,	true,	false,	false,
			true,	true,	false,	false,
			false,	false,	false,	false,
		}}),
	)
})

var _ = Describe("random 통제 테스트", func() {
	It("seed 값이 같으면 동일한 결과가 나온다.", func() {
		rand.Seed(0)
		expected := randomMino()

		rand.Seed(0)
		actual := randomMino()
		diff := deep.Equal(expected.cells, actual.cells)
		Expect(diff).Should(BeNil())
	})

	It("seed 값이 다르면 결과도 다르게.", func() {
		rand.Seed(0)
		expected := randomMino()

		rand.Seed(1)
		actual := randomMino()
		diff := deep.Equal(expected.cells, actual.cells)
		Expect(diff).ShouldNot(BeNil())
	})
})

var _ = Describe("mino.RenderInfo() 함수가", func() {
	It("렌더링 정보를 제대로 반환한다.", func() {
		m := mino{
			cells: 		[][]cell{{false, true, false, true}, {true, true, true, true}, {true, true, false, true}, {false, true, true, true}},
			color: 		123,
			x:     		1234,
			y:     		4321,
			rotation:	0,
		}
		expected := []renderer.Info{
			&render.InfoImpl{PosX: 1 + 1234, PosY: 4321, Color: 123},
			&render.InfoImpl{PosX: 3 + 1234, PosY: 4321, Color: 123},
		}
		actual := m.RenderInfo()
		Expect(actual).Should(Equal(expected))
	})
})
