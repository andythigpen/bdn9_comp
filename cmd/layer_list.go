package cmd

import (
	"fmt"

	"github.com/andythigpen/bdn9_comp/v2/serial"
	"github.com/spf13/cobra"
)

// layerListCmd represents the layer set command
var layerListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the keyboard layers",
	RunE: func(cmd *cobra.Command, args []string) error {
		var i serial.Layer
		for i = serial.LAYER_DEFAULT; i < serial.LAYER_MAX; i++ {
			fmt.Println(i.String())
		}
		return nil
	},
}

func init() {
	layerCmd.AddCommand(layerListCmd)
}
