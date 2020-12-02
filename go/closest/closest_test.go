package closest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ClosestToZeroTestCase struct {
	name string
	numbers []int
	expected int
}

var TestCases = []ClosestToZeroTestCase {
	{name:"Positive Integers", numbers:[]int{4, 3, 2, 1}, expected:1},
	{name:"Mixed Integers", numbers:[]int{-1, 2, 3}, expected:-1},
	{name:"Zero", numbers:[]int{0, 1}, expected:0},
	{name:"Several negative Integers", numbers:[]int{-2, -3, -1}, expected:-1},
	{name:"Empty list", numbers:[]int{}, expected:0},
}

func TestClosestToZero(t *testing.T) {
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, findClosestToZero(tc.numbers))
		})
	}
}

