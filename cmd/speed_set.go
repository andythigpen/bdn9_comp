package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var speed uint8

// speedSetCmd represents the speed set command
var speedSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the RGB matrix animation speed",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var s uint64
		if s, err = strconv.ParseUint(args[0], 10, 8); err != nil {
			return err
		}
		return device.SetSpeed(uint8(s))
	},
}

func init() {
	speedCmd.AddCommand(speedSetCmd)
}
