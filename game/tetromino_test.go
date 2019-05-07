package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"lets-go-tetris/render"
	"math/rand"
)

var _ = Describe("tetromino initialize test", func() {
	type testData struct {
		input    []string
		expected []cell
	}

	DescribeTable("test cases", func(d testData) {
		m := tetromino{}
		m.init(d.input)
		actual := m.currentCells()
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("S", testData{shapes[S], []cell{
			x, o, o, x,
			o, o, x, x,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("Z", testData{shapes[Z], []cell{
			o, o, x, x,
			x, o, o, x,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("T", testData{shapes[T], []cell{
			x, o, x, x,
			o, o, o, x,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("I", testData{shapes[I], []cell{
			x, x, x, x,
			o, o, o, o,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("O", testData{shapes[O], []cell{
			x, o, o, x,
			x, o, o, x,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("L", testData{shapes[L], []cell{
			x, x, o, x,
			o, o, o, x,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("J", testData{shapes[J], []cell{
			o, x, x, x,
			o, o, o, x,
			x, x, x, x,
			x, x, x, x,
		}}),
	)
})

var _ = Describe("newTetromino test", func() {
	It("Succes return s mino", func() {
		m := newTetromino(S)
		expected := []cell{
			x, o, o, x,
			o, o, x, x,
			x, x, x, x,
			x, x, x, x,
		}

		actual := m.currentCells()
		diff := deep.Equal(actual, expected)
		Expect(diff).Should(BeNil())
	})
})

var _ = Describe("rotate test", func() {
	type testData struct {
		shape    Shape
		rotation Rotation
		expected []cell
	}

	DescribeTable("test cases", func(d testData) {
		m := newTetromino(d.shape)
		m.rotate(d.rotation)
		actual := m.currentCells()
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("S", testData{S, -1, []cell{
			o, x, x, x,
			o, o, x, x,
			x, o, x, x,
			x, x, x, x,
		}}),
		Entry("Z", testData{Z, 6, []cell{
			x, x, x, x,
			o, o, x, x,
			x, o, o, x,
			x, x, x, x,
		}}),
		Entry("T", testData{T, -2, []cell{
			x, x, x, x,
			o, o, o, x,
			x, o, x, x,
			x, x, x, x,
		}}),
		Entry("I", testData{I, 16, []cell{
			x, x, x, x,
			o, o, o, o,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("O", testData{O, -4, []cell{
			x, o, o, x,
			x, o, o, x,
			x, x, x, x,
			x, x, x, x,
		}}),
		Entry("L", testData{L, 7, []cell{
			o, o, x, x,
			x, o, x, x,
			x, o, x, x,
			x, x, x, x,
		}}),
		Entry("J", testData{J, -17, []cell{
			x, o, x, x,
			x, o, x, x,
			o, o, x, x,
			x, x, x, x,
		}}),
	)
})

var _ = Describe("rotateClockWise test", func() {
	type testData struct {
		shape    Shape
		rotation Rotation
		expected Rotate
	}

	DescribeTable("test cases", func(d testData) {
		m := newTetromino(d.shape)
		m.rotate(d.rotation)
		actual := m.rotateClockWise()
		Expect(actual).Should(Equal(d.expected))
	},
		Entry("S", testData{S, Zero, ZtoR}),
		Entry("L", testData{L, Right, RtoT}),
		Entry("O", testData{O, Two, TtoL}),
		Entry("Z", testData{Z, Left, LtoZ}),
	)
})

var _ = Describe("rotateCounterClockWise test", func() {
	type testData struct {
		shape    Shape
		rotation Rotation
		expected Rotate
	}

	DescribeTable("test cases", func(d testData) {
		m := newTetromino(d.shape)
		m.rotate(d.rotation)
		actual := m.rotateCounterClockWise()
		Expect(actual).Should(Equal(d.expected))
	},
		Entry("S", testData{S, Zero, ZtoL}),
		Entry("L", testData{L, Left, LtoT}),
		Entry("O", testData{O, Two, TtoR}),
		Entry("Z", testData{Z, Right, RtoZ}),
	)
})

var _ = Describe("wallKick test", func() {
	type testData struct {
		shape          Shape
		x, y           int
		rotate         Rotate
		rotation       Rotation
		width          int
		height         int
		ground         []cell
		expectedX      int
		expectedY      int
		expectedReturn bool
	}

	DescribeTable("test cases", func(d testData) {
		g := ground{width: d.width, height: d.height}
		g.cells = d.ground

		actual := newTetromino(d.shape)
		actual.x = d.x
		actual.y = d.y
		actual.rotate(d.rotation)
		actualReturn := actual.wallKick(&g, d.rotate)

		Expect(actualReturn).Should(Equal(d.expectedReturn))
		Expect(actual.x).Should(Equal(d.expectedX))
		Expect(actual.y).Should(Equal(d.expectedY))
	},
		Entry("I", testData{I, 1, 3, LtoT, Two, 10, 8, []cell{
			x, x, x, x, x, x, x, x, x, x,
			x, x, x, x, x, x, x, x, x, x,
			x, x, x, x, x, x, x, x, x, x,
			x, x, x, x, x, x, x, x, x, x,
			o, o, x, o, o, o, o, o, o, o,
			o, o, x, o, o, o, o, o, o, o,
			o, o, x, o, o, o, o, o, o, o,
			o, o, x, o, o, o, o, o, o, o,
		}, 2, 1, true}),
		Entry("I", testData{I, 1, 3, LtoT, Two, 10, 8, []cell{
			x, x, x, x, x, x, x, x, x, x,
			x, x, x, x, x, x, x, x, x, x,
			x, x, x, x, x, x, x, x, x, x,
			o, o, x, o, o, o, o, o, o, o,
			o, o, x, o, o, o, o, o, o, o,
			o, o, x, o, o, o, o, o, o, o,
			o, o, x, o, o, o, o, o, o, o,
			o, o, x, o, o, o, o, o, o, o,
		}, 1, 3, false}),
		Entry("J", testData{J, 3, 2, ZtoL, Left, 10, 8, []cell{
			x, x, x, x, x, x, x, x, x, x,
			x, x, x, x, o, o, x, x, x, x,
			x, x, x, x, x, o, o, o, x, x,
			x, x, x, x, x, x, o, o, o, o,
			x, o, o, o, x, x, x, o, o, o,
			o, o, x, x, x, x, o, o, o, o,
			o, o, o, o, x, x, o, o, o, o,
			o, o, o, o, o, x, o, o, o, o,
		}, 4, 4, true}),
	)
})

var _ = Describe("random 통제 테스트", func() {
	It("seed 값이 같으면 동일한 결과가 나온다.", func() {
		rand.Seed(0)
		expected := randomTetromino()

		rand.Seed(0)
		actual := randomTetromino()
		diff := deep.Equal(expected.cells, actual.cells)
		Expect(diff).Should(BeNil())
	})

	It("seed 값이 다르면 결과도 다르게.", func() {
		rand.Seed(0)
		expected := randomTetromino()

		rand.Seed(1)
		actual := randomTetromino()
		diff := deep.Equal(expected.cells, actual.cells)
		Expect(diff).ShouldNot(BeNil())
	})
})

var _ = Describe("tetromino.RenderInfo() 함수가", func() {
	It("렌더링 정보를 제대로 반환한다.", func() {
		m := tetromino{
			cells:    [][]cell{{x, o, x, o}, {o, o, o, o}, {o, o, x, o}, {x, o, o, o}},
			color:    123,
			x:        1234,
			y:        4321,
			rotation: 0,
		}
		expected := []render.Info{
			&render.InfoImpl{PosX: 1 + 1234, PosY: 4321, Color: 123},
			&render.InfoImpl{PosX: 3 + 1234, PosY: 4321, Color: 123},
		}
		actual := m.RenderInfo()
		Expect(actual).Should(Equal(expected))
	})
})
