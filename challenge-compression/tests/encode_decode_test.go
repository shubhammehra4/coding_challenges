//go:build integration

package tests

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shubhammehra4/coding_challenges/challenge-compression/cmd"
)

func TestEncodeDecode(t *testing.T) {
	testDir := t.TempDir()

	orginalFile := "135-0.txt"
	testFile := filepath.Join(testDir, "test1.txt")
	targetEncodedFile := filepath.Join(testDir, "encoded.txt")
	targetDecodedFile := filepath.Join(testDir, "decoded.txt")

	// Copy the test files to the temporary directory
	err := copyFile(orginalFile, testFile)
	require.NoError(t, err)

	rootCmd := cmd.NewRootCmd(context.TODO())
	rootCmd.SetArgs([]string{testFile, "-o", targetEncodedFile})
	err = rootCmd.Execute()
	require.NoError(t, err)
	assert.FileExists(t, targetEncodedFile)

	rootCmd = cmd.NewRootCmd(context.TODO())
	rootCmd.SetArgs([]string{targetEncodedFile, "-d", "-o", targetDecodedFile})
	err = rootCmd.Execute()
	require.NoError(t, err)
	assert.FileExists(t, targetDecodedFile)

	// Compare the original file with the decoded file
	assertFilesEqual(t, orginalFile, targetDecodedFile)
}

func assertFilesEqual(t *testing.T, file1, file2 string) {
	file1Data, err := os.ReadFile(file1)
	require.NoError(t, err)

	file2Data, err := os.ReadFile(file2)
	require.NoError(t, err)

	assert.Equal(t, file1Data, file2Data)
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
