package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmilyFace(t *testing.T) {
	testCase := []struct {
		Request  []string
		Expected int
	}{
		{[]string{":)", ";(", ";}", ":-D"}, 2},
		{[]string{";D", ":-(", ":-)", ";~)"}, 3},
		{[]string{";]", ":[", ";*", ":$", ";-D"}, 1},
	}

	for _, tc := range testCase {
		actual := CountSmilyFace(tc.Request)
		assert.Equal(t, tc.Expected, actual)
	}
}
