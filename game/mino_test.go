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

var _ = Describe("newMino 테스트", func() {
	It("S블럭을 잘 반환한다.", func() {
		m := newMino(S)
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

var _ = Describe("mino rotation 테스트", func() {
	type testData struct {
		shape    Shape
		rotation Rotation
		expected []cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		m := newMino(d.shape)
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

var _ = XDescribe("mino srs 테스트 (super rotation system)", func() {
	type testData struct {
		shape Shape
		x, y int32
		rotate Rotate
		ground []cell
		expectedX int32
		expectedY int32
	}

	DescribeTable("테스트 케이스", func(d testData) {
		g := ground{}
		g.cells = d.ground

		actual := newMino(d.shape)
		actual.srs(&g, d.rotate)

		expected := newMino(d.shape)
		expected.x = d.expectedX
		expected.y = d.expectedY

		diff := deep.Equal(actual, expected)
		Expect(diff).Should(BeNil())
	},
		Entry("J", testData{L, 4, 3, -1, []cell{
			x, x, x, x, x, x, x, x, x, x,
			x, x, x, x, o, o, x, x, x, x,
			x, x, x, x, x, o, o, o, x, x,
			x, x, x, x, x, x, o, o, o, o,
			x, o, o, o, x, x, x, o, o, o,
			o, o, x, x, x, x, o, o, o, o,
			o, o, o, o, x, x, o, o, o, o,
			o, o, o, o, o, x, o, o, o, o,
		}, 6, 2,}),
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
			cells:    [][]cell{{x, o, x, o}, {o, o, o, o}, {o, o, x, o}, {x, o, o, o}},
			color:    123,
			x:        1234,
			y:        4321,
			rotation: 0,
		}
		expected := []renderer.Info{
			&render.InfoImpl{PosX: 1 + 1234, PosY: 4321, Color: 123},
			&render.InfoImpl{PosX: 3 + 1234, PosY: 4321, Color: 123},
		}
		actual := m.RenderInfo()
		Expect(actual).Should(Equal(expected))
	})
})
