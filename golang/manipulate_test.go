package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManipulate(t *testing.T) {
	testCase := []struct {
		Request  string
		Expected []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, tc := range testCase {
		actual := Manipulate(tc.Request)
		assert.Equal(t, tc.Expected, actual)
	}
}
