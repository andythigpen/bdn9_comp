package cmd

import (
	"fmt"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// layerListCmd represents the layer set command
var layerListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the keyboard layers",
	RunE: func(cmd *cobra.Command, args []string) error {
		var i pb.Layer
		for i = pb.Layer_LAYER_DEFAULT; i < pb.Layer_LAYER_MAX; i++ {
			fmt.Println(i.Description())
		}
		return nil
	},
}

func init() {
	layerCmd.AddCommand(layerListCmd)
}
