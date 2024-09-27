package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/shubhammehra4/coding_challenges/challenge-compression/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		fmt.Println("Got signal to stop the program, please wait while we clean up!")
		cancel()
	}()

	rootCmd := cmd.NewRootCmd(ctx)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
