package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// hsvSetCmd represents the hsv set command
var hsvSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the matrix HSV values",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, s, v, err := parseHSV(args)
		if err != nil {
			return err
		}
		ctx := context.Background()
		// shortcut to set mode at same time as hsv
		if mode != uint8(pb.RGBMode_RGB_MATRIX_INVALID) {
			_, err = client.SetRGBMode(ctx, &pb.SetRGBModeRequest{Mode: uint32(mode)})
			if err != nil {
				return err
			}
		}
		if speed > 0 {
			_, err = client.SetSpeed(ctx, &pb.SetSpeedRequest{Speed: uint32(speed)})
			if err != nil {
				return err
			}
		}
		_, err = client.SetMatrixHSV(ctx, &pb.SetMatrixHSVRequest{H: uint32(h), S: uint32(s), V: uint32(v)})
		return err
	},
}

func init() {
	hsvSetCmd.Flags().Uint8VarP(&mode, "mode", "m", uint8(pb.RGBMode_RGB_MATRIX_INVALID), "Optional mode")
	hsvSetCmd.Flags().Uint8VarP(&speed, "speed", "s", 0, "Optional speed")
	hsvCmd.AddCommand(hsvSetCmd)
}
