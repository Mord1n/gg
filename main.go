package main

import (
	"gege/cmd" // Import your commands
	helloext "github.com/Mord1n/hello-ext/cmd"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gg",
		Short: "A simple CLI tool",
	}

	// Add the commands
	rootCmd.AddCommand(cmd.HelloCmd)
	rootCmd.AddCommand(helloext.HelloCmd)
	rootCmd.AddCommand(cmd.UpdateCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
