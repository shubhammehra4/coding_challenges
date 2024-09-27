package core

import (
	"os"

	"github.com/shubhammehra4/coding_challenges/challenge-wc/utils"
)

func WordCount(options WordCountOptions) (res []*WordCountResult, err error) {
	if options.stdin != nil {
		s, err := processBasedOnStrategy(options.stdin, options)
		if err != nil {
			return res, err
		}
		res = append(res, s)
		return res, nil
	}

	total := NewWordCountResult("total")
	for _, filePath := range options.filePaths {
		f, err := processFile(filePath, options)
		if err != nil {
			return nil, err
		}
		res = append(res, f)
		total.Add(f)
	}

	if len(options.filePaths) > 1 {
		res = append(res, total)
	}

	return res, nil
}

func processFile(filePath string, options WordCountOptions) (res *WordCountResult, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return processBasedOnStrategy(file, options)
}

func processBasedOnStrategy(file *os.File, options WordCountOptions) (res *WordCountResult, err error) {
	switch options.strategy {
	case utils.ChunkedStrategy:
		return calculateWordCountChunked(file, options)
	default:
		return calculateWordCount(file, options)
	}
}
