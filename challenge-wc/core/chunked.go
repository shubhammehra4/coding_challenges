package core

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"unicode"

	"github.com/shubhammehra4/coding_challenges/challenge-wc/utils"
)

func calculateWordCountChunked(f *os.File, options WordCountOptions) (res *WordCountResult, err error) {
	res = NewWordCountResultWithOptions(f.Name(), &options)
	reader := bufio.NewReader(f)
	buf := make([]byte, options.chunkSize)
	var leftover []byte

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return res, err
		}
		if n == 0 {
			break
		}
		chunk := buf[:n]
		wordChunk := append(leftover, chunk...)

		chunkRes := NewEmptyWordCountResult("")
		if options.lines {
			chunkRes.Lines = utils.IntPointer(countLines(chunk))
		}
		if options.words {
			chunkRes.Words = utils.IntPointer(countWords(wordChunk, false))
		}
		if options.bytes && !options.characters {
			chunkRes.BytesOrCharacters = utils.IntPointer(countBytes(chunk))
		}
		if options.characters {
			chunkRes.BytesOrCharacters = utils.IntPointer(countCharacters(chunk))
		}

		lastWordStart := bytes.LastIndexFunc(wordChunk, unicode.IsSpace)
		if lastWordStart != -1 {
			leftover = wordChunk[lastWordStart+1:]
		} else {
			leftover = wordChunk
		}

		res.Add(chunkRes)
	}
	return res, nil
}
