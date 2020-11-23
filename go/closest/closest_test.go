package closest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_positive_integers(t *testing.T) {
	numbers := []int{4, 3, 2, 1}
	assert.Equal(t, 1, findClosestToZero(numbers))
}

func Test_negative(t *testing.T) {
	numbers := []int{-1, 2, 3}
	assert.Equal(t, -1, findClosestToZero(numbers))
}

func Test_zero(t *testing.T) {
	numbers := []int{0, 1}
	assert.Equal(t, 0, findClosestToZero(numbers))
}
