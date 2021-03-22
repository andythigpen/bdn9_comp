package cmd

import (
	"github.com/spf13/cobra"
)

// keyOffCmd represents the off command for individual keys
var keyOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Disable indicator for specific key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return device.DisableIndicator(key)
	},
}

func init() {
	keyOffCmd.Flags().Uint8VarP(&key, "key", "k", 0, "Key index (max 11)")
	keyCmd.AddCommand(keyOffCmd)
}
