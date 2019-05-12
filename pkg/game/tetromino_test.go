package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"math/rand"
)

var _ = Describe("tetromino initialize test", func() {
	type testData struct {
		input    []string
		expected [][]Cell
	}

	DescribeTable("test cases", func(d testData) {
		m := tetromino{}
		m.init(d.input)
		actual := m.GetCells()
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("S", testData{shapes[S], [][]Cell{
			{x, o, o, x},
			{o, o, x, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("Z", testData{shapes[Z], [][]Cell{
			{o, o, x, x},
			{x, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("T", testData{shapes[T], [][]Cell{
			{x, o, x, x},
			{o, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("I", testData{shapes[I], [][]Cell{
			{x, x, x, x},
			{o, o, o, o},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("O", testData{shapes[O], [][]Cell{
			{x, o, o, x},
			{x, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("L", testData{shapes[L], [][]Cell{
			{x, x, o, x},
			{o, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("J", testData{shapes[J], [][]Cell{
			{o, x, x, x},
			{o, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
	)
})

var _ = Describe("newTetromino test", func() {
	It("Succes return s mino", func() {
		m := newTetromino(S)
		expected := [][]Cell{
			{x, o, o, x},
			{o, o, x, x},
			{x, x, x, x},
			{x, x, x, x},
		}

		actual := m.GetCells()
		diff := deep.Equal(actual, expected)
		Expect(diff).Should(BeNil())
	})
})

var _ = Describe("rotate test", func() {
	type testData struct {
		shape    Shape
		rotation Rotation
		expected [][]Cell
	}

	DescribeTable("test cases", func(d testData) {
		m := newTetromino(d.shape)
		m.rotate(d.rotation)
		actual := m.GetCells()
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("S", testData{S, -1, [][]Cell{
			{o, x, x, x},
			{o, o, x, x},
			{x, o, x, x},
			{x, x, x, x},
		}}),
		Entry("Z", testData{Z, 6, [][]Cell{
			{x, x, x, x},
			{o, o, x, x},
			{x, o, o, x},
			{x, x, x, x},
		}}),
		Entry("T", testData{T, -2, [][]Cell{
			{x, x, x, x},
			{o, o, o, x},
			{x, o, x, x},
			{x, x, x, x},
		}}),
		Entry("I", testData{I, 16, [][]Cell{
			{x, x, x, x},
			{o, o, o, o},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("O", testData{O, -4, [][]Cell{
			{x, o, o, x},
			{x, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("L", testData{L, 7, [][]Cell{
			{o, o, x, x},
			{x, o, x, x},
			{x, o, x, x},
			{x, x, x, x},
		}}),
		Entry("J", testData{J, -17, [][]Cell{
			{x, o, x, x},
			{x, o, x, x},
			{o, o, x, x},
			{x, x, x, x},
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
		Entry("S", testData{S, ZeroRotation, ZtoR}),
		Entry("L", testData{L, RightRotation, RtoT}),
		Entry("O", testData{O, TwoRotation, TtoL}),
		Entry("Z", testData{Z, LeftRotation, LtoZ}),
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
		Entry("S", testData{S, ZeroRotation, ZtoL}),
		Entry("L", testData{L, LeftRotation, LtoT}),
		Entry("O", testData{O, TwoRotation, TtoR}),
		Entry("Z", testData{Z, RightRotation, RtoZ}),
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
		matrix         [][]Cell
		expectedX      int
		expectedY      int
		expectedReturn bool
	}

	DescribeTable("test cases", func(d testData) {
		g := matrix{width: d.width, height: d.height}
		g.cells = d.matrix

		actual := newTetromino(d.shape)
		actual.x = d.x
		actual.y = d.y
		actual.rotate(d.rotation)
		actualReturn := actual.wallKick(&g, d.rotate)

		Expect(actualReturn).Should(Equal(d.expectedReturn))
		Expect(actual.x).Should(Equal(d.expectedX))
		Expect(actual.y).Should(Equal(d.expectedY))
	},
		Entry("I", testData{I, 1, 3, LtoT, TwoRotation, 10, 8, [][]Cell{
			{x, x, x, x, x, x, x, x, x, x},
			{x, x, x, x, x, x, x, x, x, x},
			{x, x, x, x, x, x, x, x, x, x},
			{x, x, x, x, x, x, x, x, x, x},
			{o, o, x, o, o, o, o, o, o, o},
			{o, o, x, o, o, o, o, o, o, o},
			{o, o, x, o, o, o, o, o, o, o},
			{o, o, x, o, o, o, o, o, o, o},
		}, 2, 1, true}),
		Entry("I", testData{I, 1, 3, LtoT, TwoRotation, 10, 8, [][]Cell{
			{x, x, x, x, x, x, x, x, x, x},
			{x, x, x, x, x, x, x, x, x, x},
			{x, x, x, x, x, x, x, x, x, x},
			{o, o, x, o, o, o, o, o, o, o},
			{o, o, x, o, o, o, o, o, o, o},
			{o, o, x, o, o, o, o, o, o, o},
			{o, o, x, o, o, o, o, o, o, o},
			{o, o, x, o, o, o, o, o, o, o},
		}, 1, 3, false}),
		Entry("J", testData{J, 3, 2, ZtoL, LeftRotation, 10, 8, [][]Cell{
			{x, x, x, x, x, x, x, x, x, x},
			{x, x, x, x, o, o, x, x, x, x},
			{x, x, x, x, x, o, o, o, x, x},
			{x, x, x, x, x, x, o, o, o, o},
			{x, o, o, o, x, x, x, o, o, o},
			{o, o, x, x, x, x, o, o, o, o},
			{o, o, o, o, x, x, o, o, o, o},
			{o, o, o, o, o, x, o, o, o, o},
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
