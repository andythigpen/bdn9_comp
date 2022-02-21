package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// keyOffCmd represents the off command for individual keys
var keyOffCmd = &cobra.Command{
	Use:   "off",
	Short: "Disable indicator for specific key",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.DisableIndicator(context.Background(), &pb.DisableIndicatorRequest{
			Key:   uint32(key),
			Layer: pb.Layer(layer),
		})
		return err
	},
}

func init() {
	keyOffCmd.Flags().Uint8VarP(&key, "key", "k", 0, "Key index (max 11)")
	keyOffCmd.Flags().Uint8VarP(&layer, "layer", "l", 0, "Layer")
	keyCmd.AddCommand(keyOffCmd)
}
