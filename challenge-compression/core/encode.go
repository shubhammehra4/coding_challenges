package core

import (
	"bytes"
	"encoding/binary"

	"github.com/shubhammehra4/coding_challenges/challenge-compression/core/huffman"
)

// encode returns the encoded prefix code table along with the encoded string
func encode(b []byte) (bytes.Buffer, error) {
	s := string(b)
	freqMap := countFrequency(s)
	root := huffman.Build(freqMap)
	table := huffman.PrefixCodeTable(root)
	return encodeContents(s, table)
}

func encodeContents(s string, table map[rune]huffman.PrefixCode) (bytes.Buffer, error) {
	var encodedText bytes.Buffer

	if err := binary.Write(&encodedText, binary.BigEndian, int32(len(table))); err != nil {
		return encodedText, err
	}
	for k, v := range table {
		if err := binary.Write(&encodedText, binary.BigEndian, int32(k)); err != nil {
			return encodedText, err
		}
		if err := binary.Write(&encodedText, binary.BigEndian, int32(v.Bits)); err != nil {
			return encodedText, err
		}
		if err := binary.Write(&encodedText, binary.BigEndian, v.Code); err != nil {
			return encodedText, err
		}
	}

	var code uint64
	var bits int32
	for _, r := range s {
		prefixCode := table[r]
		code = code<<uint(prefixCode.Bits) | prefixCode.Code
		bits += prefixCode.Bits
		for bits >= 8 {
			if err := encodedText.WriteByte(byte(code >> uint(bits-8))); err != nil {
				return encodedText, err
			}
			bits -= 8
		}
	}
	if bits > 0 {
		if err := encodedText.WriteByte(byte(code << uint(8-bits))); err != nil {
			return encodedText, err
		}
	}
	return encodedText, nil
}

func countFrequency(s string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, r := range s {
		freqMap[r]++
	}
	return freqMap
}
