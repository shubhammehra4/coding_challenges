//go:build unit

package huffman

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		name                    string
		freqMap                 map[rune]int
		expectedPrefixCodeTable map[rune]PrefixCode
	}{
		{
			name: "Basic Test Case",
			freqMap: map[rune]int{
				'A': 5,
				'B': 9,
				'C': 12,
				'D': 13,
			},
			expectedPrefixCodeTable: map[rune]PrefixCode{
				'A': {Code: 0b00, Bits: 2},
				'B': {Code: 0b01, Bits: 2},
				'C': {Code: 0b10, Bits: 2},
				'D': {Code: 0b11, Bits: 2},
			},
		},
		{
			name: "More Complex Test Case",
			freqMap: map[rune]int{
				'C': 32,
				'D': 42,
				'E': 120,
				'K': 7,
				'L': 42,
				'M': 24,
				'U': 37,
				'Z': 2,
			},
			expectedPrefixCodeTable: map[rune]PrefixCode{
				'C': {Code: 0b1110, Bits: 4},
				'D': {Code: 0b101, Bits: 3},
				'E': {Code: 0b0, Bits: 1},
				'K': {Code: 0b111101, Bits: 6},
				'L': {Code: 0b110, Bits: 3},
				'M': {Code: 0b11111, Bits: 5},
				'U': {Code: 0b100, Bits: 3},
				'Z': {Code: 0b111100, Bits: 6},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := Build(tc.freqMap)
			require.NotEqual(t, nil, root)
			table := PrefixCodeTable(root)
			require.Equal(t, tc.expectedPrefixCodeTable, table)
		})
	}
}
