package game

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test cases for ghost init method", func() {
	type testData struct {
		shape         Shape
		width, height int
	}

	type expectedData struct {
		x, y int
	}

	DescribeTable(
		"Test cases",
		func(d testData, e expectedData) {
			t := newTetromino(d.shape)
			m := newMatrix(d.width, d.height)

			g := &ghost{}
			g.init(m, t)

			x, y := g.GetPosition()
			Expect(x).Should(Equal(e.x))
			Expect(y).Should(Equal(e.y))
		},
		Entry("L",
			testData{L, 4, 4},
			expectedData{0, 2}),
	)
})
