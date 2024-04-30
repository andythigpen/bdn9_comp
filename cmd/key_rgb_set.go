package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

var key uint8
var layer uint8
var turnOn bool

// keyRgbSetCmd represents the RGB set command for individual keys
var keyRgbSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set RGB for an individual key",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, g, b, err := parseRGB(args)
		if err != nil {
			return err
		}
		pbLayer := pb.Layer(layer)
		_, err = client.SetIndicatorRGB(context.Background(), &pb.SetIndicatorRGBRequest{
			Key:   uint32(key),
			R:     uint32(r),
			G:     uint32(g),
			B:     uint32(b),
			Layer: pbLayer,
		})
		if err != nil {
			return err
		}
		if turnOn {
			_, err = client.EnableIndicator(context.Background(), &pb.EnableIndicatorRequest{
				Key:   uint32(key),
				Layer: pbLayer,
			})
			return err
		}
		return nil
	},
}

func init() {
	keyRgbSetCmd.Flags().Uint8VarP(&key, "key", "k", 0, "Key index (max 11)")
	keyRgbSetCmd.Flags().Uint8VarP(&layer, "layer", "l", 0, "Layer")
	keyRgbSetCmd.Flags().BoolVar(&turnOn, "on", false, "Also turn the key on")
	keyCmd.AddCommand(keyRgbSetCmd)
}
