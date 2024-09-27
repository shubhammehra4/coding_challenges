package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Process(co *CompressOptions) error {
	f, err := os.Open(co.filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	outputPath := co.outputPath
	if outputPath == "" {
		outputPath = filepath.Join(filepath.Dir(co.filePath), fmt.Sprintf("%s_%s", co.mode, filepath.Base(co.filePath)))
	}
	if err = os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}

	var result []byte

	switch co.mode {
	case ENCODE:
		b, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		encodedBytes, err := encode(b)
		if err != nil {
			return err
		}
		result = encodedBytes.Bytes()
		if co.showStats {
			fmt.Printf("Original size: %d bytes\n", len(b))
			fmt.Printf("Compressed size: %d bytes\n", len(result))
			fmt.Printf("Compression ratio: %.2f%%\n", float64(len(result))/float64(len(b))*100)
		}
	case DECODE:
		decodeBytes, err := decode(f)
		if err != nil {
			return err
		}
		result = decodeBytes
	default:
		return fmt.Errorf("unknown mode: %s", co.mode)
	}

	return os.WriteFile(outputPath, result, 0o644)
}
