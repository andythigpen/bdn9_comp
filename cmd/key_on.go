package cmd

import (
	"github.com/spf13/cobra"
)

// keyOnCmd represents the on command for individual keys
var keyOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Enable indicator for specific key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return device.EnableIndicator(key)
	},
}

func init() {
	keyOnCmd.Flags().Uint8VarP(&key, "key", "k", 0, "Key index (max 11)")
	keyCmd.AddCommand(keyOnCmd)
}
