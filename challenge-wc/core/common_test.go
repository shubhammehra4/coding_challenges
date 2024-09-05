package core

import (
	"testing"

	"github.com/shubhammehra4/coding_challenges/challenge-wc/utils"
)

func TestCountLines(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "Empty buffer",
			input:    []byte{},
			expected: 0,
		},
		{
			name:     "Single line",
			input:    []byte("line1"),
			expected: 0,
		},
		{
			name:     "Buffer with only newlines",
			input:    []byte(string(utils.EndOfLine) + string(utils.EndOfLine) + string(utils.EndOfLine)),
			expected: 3,
		},
		{
			name:     "Line ends with newline",
			input:    []byte("line1" + string(utils.EndOfLine)),
			expected: 1,
		},
		{
			name:     "Multiple consecutive newlines",
			input:    []byte("line1" + string(utils.EndOfLine) + string(utils.EndOfLine) + "line2" + string(utils.EndOfLine)),
			expected: 3,
		},
		{
			name:     "Line with special characters",
			input:    []byte("line1" + string(utils.EndOfLine) + "li\tne2" + string(utils.EndOfLine) + "line3"),
			expected: 2,
		},
		{
			name:     "Line with non-ASCII characters",
			input:    []byte("line1" + string(utils.EndOfLine) + "liñe2" + string(utils.EndOfLine) + "line3"),
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countLines(tt.input)
			if result != tt.expected {
				t.Errorf("countLines() = %d; want %d", result, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "Empty buffer",
			input:    []byte{},
			expected: 0,
		},
		{
			name:     "Single word",
			input:    []byte("word"),
			expected: 1,
		},
		{
			name:     "Multiple words with multiple spaces",
			input:    []byte("word1  word2   word3"),
			expected: 3,
		},
		{
			name:     "Words separated by newline",
			input:    []byte("word1\nword2\nword3"),
			expected: 3,
		},
		{
			name:     "Words separated by tabs",
			input:    []byte("word1\tword2\tword3"),
			expected: 3,
		},
		{
			name:     "Leading and trailing spaces",
			input:    []byte("   word1  word2   "),
			expected: 2,
		},
		{
			name:     "Spaces only",
			input:    []byte("     "),
			expected: 0,
		},
		{
			name:     "Punctuation between words",
			input:    []byte("word1,word2.word3"),
			expected: 1,
		},
		{
			name:     "Unicode characters",
			input:    []byte("word1 αβγ word2"),
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countWords(tt.input, true)
			if result != tt.expected {
				t.Errorf("countWords() = %d; want %d", result, tt.expected)
			}
		})
	}
}
