package cmd

import (
	"github.com/spf13/cobra"
)

// speedCmd represents the speed command
var speedCmd = &cobra.Command{
	Use:   "speed",
	Short: "RGB matrix animation speed",
}

func init() {
	rootCmd.AddCommand(speedCmd)
}
