package cmd

import (
	"context"
	"strconv"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// keyOffCmd represents the off command for individual keys
var keyOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Disable indicator for specific key(s)",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var k uint64
		var err error
		ctx := context.Background()
		for _, arg := range args {
			if k, err = strconv.ParseUint(arg, 10, 8); err != nil {
				return err
			}
			_, err = client.DisableIndicator(ctx, &pb.DisableIndicatorRequest{
				Key:   uint32(k),
				Layer: pb.Layer(layer),
			})
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	keyOffCmd.Flags().Uint8VarP(&layer, "layer", "l", 0, "Layer")
	keyCmd.AddCommand(keyOffCmd)
}
