package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test of matrix reset", func() {
	type testData struct {
		x, y     int
		expected [][]Cell
	}

	DescribeTable("Test cases", func(d testData) {
		m := newMatrix(d.x, d.y)
		actual := m.Cells()
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("2x3", testData{2, 3, [][]Cell{
			{x, x},
			{x, x},
			{x, x},
		}}),
		Entry("4x3", testData{4, 3, [][]Cell{
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
	)
})

var _ = Describe("Test of matrix merge", func() {
	type testData struct {
		s        Shape
		x, y     int
		expected [][]Cell
	}

	m := newMatrix(4, 10)

	DescribeTable("Test cases", func(d testData) {
		t := tetromino{x: d.x, y: d.y}
		t.init(shapes[d.s])
		m.merge(&t)

		actual := m.Cells()
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("L", testData{L, 0, 0, [][]Cell{
			{x, x, o, x},
			{o, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("I", testData{I, 2, 3, [][]Cell{
			{x, x, o, x},
			{o, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, o, o},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
		Entry("O", testData{O, 3, 5, [][]Cell{
			{x, x, o, x},
			{o, o, o, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, o, o},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
		}}),
	)
})

var _ = Describe("Test of matrix collide", func() {
	It("Case of not collide", func() {
		m := newMatrix(4, 10)

		t := newTetromino(S)
		m.merge(t)

		t = newTetromino(L)
		t.y = 4

		actual := m.Collide(t)

		Expect(actual).Should(Equal(false))
	})

	It("Case of collide with already merged block", func() {
		m := newMatrix(4, 10)

		t := newTetromino(S)
		m.merge(t)

		t = newTetromino(L)
		actual := m.Collide(t)

		Expect(actual).Should(Equal(true))
	})

	It("Case of out of range matrix", func() {
		m := newMatrix(4, 10)

		t := newTetromino(I)
		t.x = 100
		t.y = 100

		actual := m.Collide(t)

		Expect(actual).Should(Equal(true))
	})
})

var _ = Describe("Test of matrix remove lines", func() {
	type testData struct {
		expected int
		before   [][]Cell
		after    [][]Cell
	}

	DescribeTable("Test cases", func(d testData) {
		m := newMatrix(4, 10)
		m.cells = d.before

		actual := m.removeLines()
		Expect(actual).Should(Equal(d.expected))

		diff := deep.Equal(m.cells, d.after)
		Expect(diff).Should(BeNil())
	},
		Entry("1 line", testData{1, [][]Cell{
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, o, x, o},
			{o, o, o, o},
			{o, o, x, o},
		}, [][]Cell{
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, o, x, o},
			{o, o, x, o},
		}}),
		Entry("3 lines", testData{3, [][]Cell{
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{o, o, o, o},
			{o, o, o, o},
			{o, o, x, o},
			{o, o, o, o},
		}, [][]Cell{
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{o, o, x, o},
		}}),
		Entry("4 lines", testData{4, [][]Cell{
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, o, x, x},
			{x, o, x, x},
			{o, o, o, o},
			{o, o, o, o},
			{o, o, o, o},
			{o, o, o, o},
		}, [][]Cell{
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, x, x, x},
			{x, o, x, x},
			{x, o, x, x},
		}}),
	)
})
