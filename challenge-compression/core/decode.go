package core

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"

	"github.com/shubhammehra4/coding_challenges/challenge-compression/core/huffman"
)

// decode reads the encoded file and decodes it
func decode(f *os.File) ([]byte, error) {
	tableSize := int32(0)
	if err := binary.Read(f, binary.BigEndian, &tableSize); err != nil {
		return nil, err
	}

	decodedTable := make(map[huffman.PrefixCode]rune)
	for i := int32(0); i < tableSize; i++ {
		var char rune
		var bits int32
		var code uint64
		if err := binary.Read(f, binary.BigEndian, &char); err != nil {
			return nil, err
		}
		if err := binary.Read(f, binary.BigEndian, &bits); err != nil {
			return nil, err
		}
		if err := binary.Read(f, binary.BigEndian, &code); err != nil {
			return nil, err
		}
		decodedTable[huffman.PrefixCode{Code: code, Bits: bits}] = char
	}

	encodedContent, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return decodeContents(encodedContent, decodedTable), nil
}

func decodeContents(encodedData []byte, table map[huffman.PrefixCode]rune) []byte {
	decodedText := make([]byte, 0, len(encodedData))
	decodedBuffer := bytes.NewBuffer(decodedText)
	var code uint64
	var bits int32
	for _, bit := range encodedData {
		bitUint64 := uint64(bit)
		for i := 0; i < 8; i++ {
			code = code<<1 | (bitUint64 >> uint(7-i) & 1)
			bits++
			if tableEntry, ok := table[huffman.PrefixCode{Code: code, Bits: bits}]; ok {
				decodedBuffer.WriteRune(tableEntry)
				code = 0
				bits = 0
			}
		}
	}
	return decodedBuffer.Bytes()
}
