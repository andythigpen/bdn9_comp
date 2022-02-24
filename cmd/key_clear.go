package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// keyClearCmd clears all indicator keys
var keyClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Disable indicators for all keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		ctx := context.Background()
		for k := 0; k < pb.KEY_MAX; k++ {
			_, err = client.DisableIndicator(ctx, &pb.DisableIndicatorRequest{
				Layer: pb.Layer(layer),
				Key:   uint32(k),
			})
			if err != nil {
				return err
			}
		}
		return err
	},
}

func init() {
	keyClearCmd.Flags().Uint8VarP(&key, "key", "k", 0, "Key index (max 11)")
	keyClearCmd.Flags().Uint8VarP(&layer, "layer", "l", 0, "Layer")
	keyCmd.AddCommand(keyClearCmd)
}
