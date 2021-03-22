package cmd

import (
	"github.com/spf13/cobra"
)

var key uint8

// keyHsvSetCmd represents the hsv set command for individual keys
var keyHsvSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set HSV for an individual key",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, s, v, err := parseHSV(args)
		if err != nil {
			return err
		}
		return device.SetIndicatorHSV(key, uint8(h), uint8(s), uint8(v))
	},
}

func init() {
	keyHsvSetCmd.Flags().Uint8VarP(&key, "key", "k", 0, "Key index (max 11)")
	keyCmd.AddCommand(keyHsvSetCmd)
}
