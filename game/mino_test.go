package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
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
