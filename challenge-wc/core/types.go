package core

import (
	"fmt"
	"os"

	"github.com/shubhammehra4/coding_challenges/challenge-wc/utils"
)

type WordCountOptions struct {
	filePaths []string
	stdin     *os.File
	// flags
	lines      bool
	words      bool
	bytes      bool
	characters bool

	strategy  utils.ReadStrategy
	chunkSize int
}

func NewWordCountOptions() WordCountOptions {
	return WordCountOptions{}
}

func (wco *WordCountOptions) WithFilePaths(filePaths []string) *WordCountOptions {
	wco.filePaths = filePaths
	return wco
}

func (wco *WordCountOptions) WithStdin(stdIn *os.File) *WordCountOptions {
	wco.stdin = stdIn
	return wco
}

func (wco *WordCountOptions) WithStrategy(strategy utils.ReadStrategy) *WordCountOptions {
	wco.strategy = strategy
	return wco
}

func (wco *WordCountOptions) WithLines(lines bool) *WordCountOptions {
	wco.lines = lines
	return wco
}

func (wco *WordCountOptions) WithWords(words bool) *WordCountOptions {
	wco.words = words
	return wco
}

func (wco *WordCountOptions) WithBytes(bytes bool) *WordCountOptions {
	wco.bytes = bytes
	return wco
}

func (wco *WordCountOptions) WithCharacters(characters bool) *WordCountOptions {
	wco.characters = characters
	return wco
}

func (wco *WordCountOptions) WithChunkSize(chunkSize int) *WordCountOptions {
	wco.chunkSize = chunkSize
	return wco
}

func (wco *WordCountOptions) SetDefaultFlagsIfNone() {
	// no flags provided, enable default flags (-wcl)
	if !wco.bytes && !wco.lines && !wco.words && !wco.characters {
		wco.bytes = true
		wco.lines = true
		wco.words = true
	}
}

type WordCountResult struct {
	filePath          string
	Lines             *int
	Words             *int
	BytesOrCharacters *int
}

func NewEmptyWordCountResult(filePath string) *WordCountResult {
	return &WordCountResult{filePath: filePath}
}

func NewWordCountResult(filePath string) *WordCountResult {
	return &WordCountResult{
		filePath:          filePath,
		Lines:             utils.IntPointer(0),
		Words:             utils.IntPointer(0),
		BytesOrCharacters: utils.IntPointer(0),
	}
}

func NewWordCountResultWithOptions(filePath string, options *WordCountOptions) *WordCountResult {
	res := NewEmptyWordCountResult(filePath)
	if options.lines {
		res.Lines = utils.IntPointer(0)
	}
	if options.words {
		res.Words = utils.IntPointer(0)
	}
	if options.bytes || options.characters {
		res.BytesOrCharacters = utils.IntPointer(0)
	}
	return res
}

func (r *WordCountResult) String() string {
	res := ""
	if r.Lines != nil {
		res += fmt.Sprintf("%10d ", *r.Lines)
	}
	if r.Words != nil {
		res += fmt.Sprintf("%10d ", *r.Words)
	}
	if r.BytesOrCharacters != nil {
		res += fmt.Sprintf("%10d ", *r.BytesOrCharacters)
	}

	return res + fmt.Sprintf("%10s", r.filePath)
}

func (r *WordCountResult) Add(other *WordCountResult) {
	if r.Lines != nil && other.Lines != nil {
		*r.Lines += *other.Lines
	}
	if r.Words != nil && other.Words != nil {
		*r.Words += *other.Words
	}
	if r.BytesOrCharacters != nil && other.BytesOrCharacters != nil {
		*r.BytesOrCharacters += *other.BytesOrCharacters
	}
}
