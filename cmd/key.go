package cmd

import (
	"github.com/spf13/cobra"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Individual indicator keys",
}

func init() {
	rootCmd.AddCommand(keyCmd)
}
