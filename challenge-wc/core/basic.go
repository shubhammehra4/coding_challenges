package core

import (
	"io"
	"os"

	"github.com/shubhammehra4/coding_challenges/challenge-wc/utils"
)

// The basic version reads the file in a single pass and counts the number of lines, words, bytes or characters based on the options provided.
func calculateWordCount(file *os.File, options WordCountOptions) (res *WordCountResult, err error) {
	buf, err := io.ReadAll(file)
	if err != nil {
		return res, err
	}
	res = NewEmptyWordCountResult(file.Name())

	if options.lines {
		res.Lines = utils.IntPointer(countLines(buf))
	}
	if options.words {
		res.Words = utils.IntPointer(countWords(buf, true))
	}
	if options.bytes && !options.characters {
		res.BytesOrCharacters = utils.IntPointer(countBytes(buf))
	}
	if options.characters {
		res.BytesOrCharacters = utils.IntPointer(countCharacters(buf))
	}
	return res, nil
}
