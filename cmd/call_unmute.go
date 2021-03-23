package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// callUnmuteCmd represents the call unmute command
var callUnmuteCmd = &cobra.Command{
	Use:   "unmute",
	Short: "Unmutes call mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.SetMuteStatus(context.Background(), &pb.SetMuteStatusRequest{Muted: false})
		return err
	},
}

func init() {
	callCmd.AddCommand(callUnmuteCmd)
}
