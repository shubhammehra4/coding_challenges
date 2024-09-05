package core

import (
	"unicode"

	"github.com/shubhammehra4/coding_challenges/challenge-wc/utils"
)

func countLines(buf []byte) int {
	lines := 0
	for _, b := range buf {
		if b == utils.EndOfLine {
			lines++
		}
	}
	return lines
}

func countWords(buf []byte, countLastWord bool) int {
	words := 0
	inWord := false
	for _, b := range buf {
		if unicode.IsSpace(rune(b)) {
			if inWord {
				words++
				inWord = false
			}
		} else {
			inWord = true
		}
	}
	if countLastWord && inWord {
		words++
	}
	return words
}

func countBytes(buf []byte) int {
	return len(buf)
}

func countCharacters(buf []byte) int {
	return len([]rune(string(buf)))
}

// alternative
// func countCharacters(buf []byte) int {
// 	var runes []rune
//     for len(b) > 0 {
//         r, size := utf8.DecodeRune(b)
//         runes = append(runes, r)
//         b = b[size:]
//     }
//     return runes
// }
