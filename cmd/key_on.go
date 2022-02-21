package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// keyOnCmd represents the on command for individual keys
var keyOnCmd = &cobra.Command{
	Use:   "on",
	Short: "Enable indicator for specific key",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.EnableIndicator(context.Background(), &pb.EnableIndicatorRequest{
			Key:   uint32(key),
			Layer: uint32(layer),
		})
		return err
	},
}

func init() {
	keyOnCmd.Flags().Uint8VarP(&key, "key", "k", 0, "Key index (max 11)")
	keyOnCmd.Flags().Uint8VarP(&layer, "layer", "l", 0, "Layer")
	keyCmd.AddCommand(keyOnCmd)
}
