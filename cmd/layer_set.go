package cmd

import (
	"strconv"

	"github.com/andythigpen/bdn9_comp/v2/serial"
	"github.com/spf13/cobra"
)

// layerSetCmd represents the layer set command
var layerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the current active layer",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := strconv.ParseUint(args[0], 10, 8)
		if err != nil {
			return err
		}
		layer := serial.Layer(l)
		if err := layer.IsValid(); err != nil {
			return err
		}
		return device.ActivateLayer(layer)
	},
}

func init() {
	layerCmd.AddCommand(layerSetCmd)
}
