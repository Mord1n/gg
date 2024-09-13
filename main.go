package main

import (
	"gege/cmd" // Import your commands
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gg",
		Short: "A simple CLI tool",
	}

	// Add the hello command
	rootCmd.AddCommand(cmd.HelloCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
