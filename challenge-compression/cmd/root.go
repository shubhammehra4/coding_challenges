package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/shubhammehra4/coding_challenges/challenge-compression/core"
	"github.com/shubhammehra4/coding_challenges/challenge-compression/utils"
)

func NewRootCmd(ctx context.Context) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           utils.CmdName,
		SilenceErrors: true,
		SilenceUsage:  true,
		Short:         "Compression Tool",
		Long:          "My version of a compression tool!",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			if len(args) > 1 {
				return fmt.Errorf("%s: too many arguments, expected 1, got %d", utils.CmdName, len(args))
			}
			filePath := args[0]
			co := core.NewCompressOptions(filePath)
			if err := setCompressOptions(cmd, co); err != nil {
				return fmt.Errorf("%s: getting flag values: %w", utils.CmdName, err)
			}
			if err := core.Process(co); err != nil {
				return fmt.Errorf("%s: %w", utils.CmdName, err)
			}
			return nil
		},
	}
	addCompressionFlags(rootCmd)
	return rootCmd
}

func setCompressOptions(cmd *cobra.Command, options *core.CompressOptions) error {
	if cmd.Flags().Changed("encode") && cmd.Flags().Changed("decode") {
		return fmt.Errorf("only one of --encode or --decode can be set")
	}
	if cmd.Flags().Changed("encode") {
		encode, err := cmd.Flags().GetBool("encode")
		if err != nil {
			return err
		}
		if encode {
			options.WithMode(core.ENCODE)
		}
	}
	if cmd.Flags().Changed("decode") {
		decode, err := cmd.Flags().GetBool("decode")
		if err != nil {
			return err
		}
		if decode {
			options.WithMode(core.DECODE)
		}
	}
	if cmd.Flags().Changed("output") {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}
		options.WithOutputPath(output)
	}
	if cmd.Flags().Changed("stats") {
		stats, err := cmd.Flags().GetBool("stats")
		if err != nil {
			return err
		}
		options.WithShowStats(stats)
	}
	return nil
}

func addCompressionFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("encode", "e", true, "encode mode")
	cmd.Flags().BoolP("decode", "d", false, "decode mode")
	cmd.Flags().StringP("output", "o", "", "output file path")
	cmd.Flags().BoolP("stats", "s", false, "show compression stats, only valid in encode mode")
}
