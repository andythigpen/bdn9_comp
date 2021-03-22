package cmd

import (
	"github.com/spf13/cobra"
)

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Focus a call layer & control mute status",
}

func init() {
	rootCmd.AddCommand(callCmd)
}
