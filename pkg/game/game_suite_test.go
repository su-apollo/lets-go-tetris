package game_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Game Suite")
}

/*
var _ = Describe("matrix step 테스트", func() {
	g := matrix{width: 4, height: 10}
	g.reset()

	m := newTetromino(I)
	m.y = 6
	g.merge(m)

	m = newTetromino(Z)
	m.y = 4

	It("블럭이 충돌하지 않으면서 머지되지 않고 한칸 내려갔다.", func() {
		Expect(g.step(m)).Should(Equal(false))

		expected := []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			o, o, o, o,
			x, x, x, x,
			x, x, x, x,
		}
		var y int
		y = 5

		actual := g.cells
		diff := deep.Equal(actual, expected)
		Expect(diff).Should(BeNil())
		Expect(m.y).Should(Equal(y))
	})

	It("블럭이 충돌하면서 머지된다.", func() {
		Expect(g.step(m)).Should(Equal(true))

		expected := []cell{
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			x, x, x, x,
			o, o, x, x,
			x, o, o, x,
			o, o, o, o,
			x, x, x, x,
			x, x, x, x,
		}

		actual := g.cells
		diff := deep.Equal(actual, expected)
		Expect(diff).Should(BeNil())
	})
})
*/
