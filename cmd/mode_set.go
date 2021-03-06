package cmd

import (
	"context"
	"strconv"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

var mode uint8

// modeSetCmd represents the mode set command
var modeSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the RGB matrix animation mode",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var m uint64
		if m, err = strconv.ParseUint(args[0], 10, 8); err != nil {
			return err
		}
		_, err = client.SetRGBMode(context.Background(), &pb.SetRGBModeRequest{Mode: uint32(m)})
		return err
	},
}

func init() {
	modeCmd.AddCommand(modeSetCmd)
}
