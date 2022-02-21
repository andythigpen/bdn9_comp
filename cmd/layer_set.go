package cmd

import (
	"context"
	"strconv"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
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
		_, err = client.ActivateLayer(context.Background(), &pb.ActivateLayerRequest{Layer: pb.Layer(l)})
		return err
	},
}

func init() {
	layerCmd.AddCommand(layerSetCmd)
}
