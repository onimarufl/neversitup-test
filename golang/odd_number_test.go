package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOddNumber(t *testing.T) {
	testCase := []struct {
		Request  []int
		Expected int
	}{
		{[]int{7}, 1},
		{[]int{0}, 1},
		{[]int{1, 1, 2}, 1},
		{[]int{0, 1, 0, 1, 0}, 3},
		{[]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, 1},
	}

	for _, tc := range testCase {
		actual := FindOddNumber(tc.Request)
		assert.Equal(t, tc.Expected, actual)
	}
}
