package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/shubhammehra4/coding_challenges/challenge-wc/core"
	"github.com/shubhammehra4/coding_challenges/challenge-wc/utils"
)

func NewRootCmd(ctx context.Context) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           utils.CmdName,
		SilenceErrors: true,
		SilenceUsage:  true,
		Short:         "Word Count CLI",
		Long:          "My version of Linux command line tool wc!",
		RunE: func(cmd *cobra.Command, args []string) error {
			wcOptions := core.NewWordCountOptions()
			if err := setWordCountOptions(cmd, &wcOptions); err != nil {
				return fmt.Errorf("%s: getting flag values: %w", utils.CmdName, err)
			}
			if len(args) == 0 {
				wcOptions.WithStdin(os.Stdin)
			} else {
				wcOptions.WithFilePaths(args)
			}
			result, err := core.WordCount(wcOptions)
			if err != nil {
				return fmt.Errorf("%s: %w", utils.CmdName, err)
			}
			for _, r := range result {
				fmt.Println(r)
			}
			return nil
		},
	}
	addWordCountFlags(rootCmd)

	return rootCmd
}

func setWordCountOptions(cmd *cobra.Command, options *core.WordCountOptions) error {
	bytes, err := cmd.Flags().GetBool("c")
	if err != nil {
		return err
	}
	options.WithBytes(bytes)

	lines, err := cmd.Flags().GetBool("l")
	if err != nil {
		return err
	}
	options.WithLines(lines)

	words, err := cmd.Flags().GetBool("w")
	if err != nil {
		return err
	}
	options.WithWords(words)

	characters, err := cmd.Flags().GetBool("m")
	if err != nil {
		return err
	}
	options.WithCharacters(characters)
	options.SetDefaultFlagsIfNone()

	strategy, err := cmd.Flags().GetString("strategy")
	if err != nil {
		return err
	}
	options.WithStrategy(utils.GetStrategy(strategy))

	chunkSize, err := cmd.Flags().GetInt("chunk-size")
	if err != nil {
		return err
	}
	options.WithChunkSize(chunkSize)
	return nil
}

func addWordCountFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("c", "c", false, "The number of bytes in each input file is written to the standard output.  This will cancel out any prior usage of the -m option.")
	cmd.Flags().BoolP("l", "l", false, "The number of lines in each input file is written to the standard output.")
	cmd.Flags().BoolP("w", "w", false, "The number of words in each input file is written to the standard output.")
	cmd.Flags().BoolP("m", "m", false, "The number of characters in each input file is written to the standard output.  If the current locale does not support multibyte characters, this is equivalent to the -c option.  This will cancel out any prior usage of the -c option.")

	// strategy
	cmd.Flags().StringP("strategy", "s", "", "The strategy to use for reading file(s). Available strategies: default, chunked")
	cmd.Flags().IntP("chunk-size", "z", utils.DefaultChunkSize, "The size of the chunk to read the file in bytes. Only applicable when strategy is chunked")
}
