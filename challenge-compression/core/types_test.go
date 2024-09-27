//go:build unit

package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompressOptions_WithOutputPath(t *testing.T) {
	tests := []struct {
		initialPath string
		outputPath  string
	}{
		{
			initialPath: "/path/to/file.txt",
			outputPath:  "/custom/path/to/compressed_file.txt",
		},
		{
			initialPath: "relative/path/to/file.txt",
			outputPath:  "relative/path/to/custom_compressed_file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.initialPath, func(t *testing.T) {
			opts := NewCompressOptions(tt.initialPath).WithOutputPath(tt.outputPath)
			require.Equal(t, tt.initialPath, opts.filePath)
			require.Equal(t, tt.outputPath, opts.outputPath)
		})
	}
}

func TestCompressOptions_WithShowStats(t *testing.T) {
	tests := []struct {
		initialPath string
		showStats   bool
	}{
		{
			initialPath: "/path/to/file.txt",
			showStats:   true,
		},
		{
			initialPath: "relative/path/to/file.txt",
			showStats:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.initialPath, func(t *testing.T) {
			opts := NewCompressOptions(tt.initialPath).WithShowStats(tt.showStats)
			require.Equal(t, tt.initialPath, opts.filePath)
			require.Equal(t, tt.showStats, opts.showStats)
		})
	}
}

func TestCompressOptions_WithMode(t *testing.T) {
	tests := []struct {
		initialPath string
		mode        MODE
	}{
		{
			initialPath: "/path/to/file.txt",
			mode:        ENCODE,
		},
		{
			initialPath: "relative/path/to/file.txt",
			mode:        DECODE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.initialPath, func(t *testing.T) {
			opts := NewCompressOptions(tt.initialPath).WithMode(tt.mode)
			require.Equal(t, tt.initialPath, opts.filePath)
			require.Equal(t, tt.mode, opts.mode)
		})
	}
}
