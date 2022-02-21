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
		_, err = client.SetMatrixHSV(context.Background(), &pb.SetMatrixHSVRequest{H: uint32(h), S: uint32(s), V: uint32(v)})
		return err
	},
}

func init() {
	hsvCmd.AddCommand(hsvSetCmd)
}
