package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

var key uint8
var layer uint8
var turnOn bool

// keyHsvSetCmd represents the hsv set command for individual keys
var keyHsvSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set HSV for an individual key",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, s, v, err := parseHSV(args)
		if err != nil {
			return err
		}
		pbLayer := pb.Layer(layer)
		_, err = client.SetIndicatorHSV(context.Background(), &pb.SetIndicatorHSVRequest{
			Key:   uint32(key),
			H:     uint32(h),
			S:     uint32(s),
			V:     uint32(v),
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
	keyHsvSetCmd.Flags().Uint8VarP(&key, "key", "k", 0, "Key index (max 11)")
	keyHsvSetCmd.Flags().Uint8VarP(&layer, "layer", "l", 0, "Layer")
	keyHsvSetCmd.Flags().BoolVar(&turnOn, "on", false, "Also turn the key on")
	keyCmd.AddCommand(keyHsvSetCmd)
}
