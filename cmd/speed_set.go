package cmd

import (
	"context"
	"strconv"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

var speed uint8

// speedSetCmd represents the speed set command
var speedSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the RGB matrix animation speed",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var s uint64
		if s, err = strconv.ParseUint(args[0], 10, 8); err != nil {
			return err
		}
		_, err = client.SetSpeed(context.Background(), &pb.SetSpeedRequest{Speed: uint32(s)})
		return err
	},
}

func init() {
	speedCmd.AddCommand(speedSetCmd)
}
