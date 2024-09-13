package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// HelloCmd represents the hello command
var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints 'hello world'",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}
