package cmd

import (
	"github.com/spf13/cobra"
)

// hsvSetCmd represents the hsv set command
var hsvSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the RGB matrix animation mode",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, s, v, err := parseHSV(args)
		if err != nil {
			return err
		}
		return device.SetMatrixHSV(uint8(h), uint8(s), uint8(v))
	},
}

func init() {
	hsvCmd.AddCommand(hsvSetCmd)
}
