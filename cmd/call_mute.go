package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// callMuteCmd represents the call mute command
var callMuteCmd = &cobra.Command{
	Use:   "mute",
	Short: "Mutes call mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.SetMuteStatus(context.Background(), &pb.SetMuteStatusRequest{Muted: true})
		return err
	},
}

func init() {
	callCmd.AddCommand(callMuteCmd)
}
