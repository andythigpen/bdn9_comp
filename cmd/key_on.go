package cmd

import (
	"context"
	"strconv"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// keyOnCmd represents the on command for individual keys
var keyOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Enable indicator for specific key(s)",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var k uint64
		var err error
		ctx := context.Background()
		for _, arg := range args {
			if k, err = strconv.ParseUint(arg, 10, 8); err != nil {
				return err
			}
			_, err := client.EnableIndicator(ctx, &pb.EnableIndicatorRequest{
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
	keyOnCmd.Flags().Uint8VarP(&layer, "layer", "l", 0, "Layer")
	keyCmd.AddCommand(keyOnCmd)
}
