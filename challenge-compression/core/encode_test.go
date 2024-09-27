//go:build unit

package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountFrequency(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{"", map[rune]int{}},
		{"a", map[rune]int{'a': 1}},
		{"ab", map[rune]int{'a': 1, 'b': 1}},
		{"aa", map[rune]int{'a': 2}},
		{"aabb", map[rune]int{'a': 2, 'b': 2}},
		{"abcabc", map[rune]int{'a': 2, 'b': 2, 'c': 2}},
	}

	for _, test := range tests {
		result := countFrequency(test.input)
		require.Equal(t, test.expected, result)
	}
}
