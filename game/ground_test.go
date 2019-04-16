package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("ground reset 테스트", func() {
	type testData struct {
		x, y		int32
		expected 	[]cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		g := ground{x: d.x, y: d.y}
		g.reset()
		actual := g.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("2x3", testData{2, 3, []cell{
			false,	false,
			false,	false,
			false,	false,
		}}),
		Entry("4x3", testData{4, 3, []cell{
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
	)
})

var _ = Describe("ground merge 테스트", func() {
	type testData struct {
		s			Shape
		x, y		int32
		expected	[]cell
	}

	g := ground{x: 4, y:10}
	g.reset()

	DescribeTable("테스트 케이스", func(d testData) {
		m := mino{x: d.x, y:d.y}
		m.init(shapes[d.s])
		g.merge(&m)

		actual := g.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("L", testData{L, 0,0, []cell{
			false,	false,	true,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("I", testData{I, 2,3, []cell{
			false,	false,	true,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	true,	true,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("O", testData{O, 3,5, []cell{
			false,	false,	true,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	true,	true,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
	)
})

var _ = XDescribe("mino, ground 충돌 테스트", func() {

})

var _ = XDescribe("mino wall kick 테스트", func() {
})

var _ = XDescribe("ground tetris check 테스트", func() {
})
