package cmd

import (
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the keyboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return device.Reset()
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
