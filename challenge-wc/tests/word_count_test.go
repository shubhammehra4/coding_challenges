//go:build integration

package tests

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shubhammehra4/coding_challenges/challenge-wc/cmd"
)

func TestWordCoundCli(t *testing.T) {
	testDir := t.TempDir()

	orginalFile := "test.txt"
	testFile := filepath.Join(testDir, "test1.txt")

	// Copy the test files to the temporary directory
	err := copyFile(orginalFile, testFile)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Test with single file",
			args:     []string{testFile},
			expected: "      7145      58164     342190 " + testFile + "\n",
		},
		{
			name:     "Test with multiple files",
			args:     []string{testFile, testFile},
			expected: "      7145      58164     342190 " + testFile + "\n" + "      7145      58164     342190 " + testFile + "\n" + "     14290     116328     684380      total\n",
		},
		{
			name:     "Test with characters flag",
			args:     []string{"-m", testFile},
			expected: "    339292 " + testFile + "\n",
		},
		{
			name:     "Test with chunked strategy",
			args:     []string{"-s", "chunked", testFile},
			expected: "      7145      58164     342190 " + testFile + "\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := executeCliAndGetOutput(t, tc.args)
			require.Equal(t, tc.expected, output)
		})
	}
}

func executeCliAndGetOutput(t *testing.T, args []string) string {
	// Save original stdout
	originalStdout := os.Stdout

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd := cmd.NewRootCmd(context.TODO())
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	require.NoError(t, err)

	w.Close()
	os.Stdout = originalStdout

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
