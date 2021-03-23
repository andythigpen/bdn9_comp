package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/andythigpen/bdn9_comp/v2/serial"
	"github.com/spf13/cobra"
)

var slack, teams bool

// callStartCmd represents the call start command
var callStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a call mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		var layer uint32
		var muted bool
		if teams {
			layer = serial.LAYER_TEAMS
			muted = true // teams starts muted
		} else {
			layer = serial.LAYER_SLACK
			muted = false // slack does not start muted
		}
		req := &pb.ActivateLayerRequest{Layer: layer}
		_, err := client.ActivateLayer(context.Background(), req)
		if err != nil {
			return err
		}
		_, err = client.SetMuteStatus(context.Background(), &pb.SetMuteStatusRequest{Muted: muted})
		return err
	},
}

func init() {
	callStartCmd.Flags().BoolVar(&slack, "slack", false, "Slack call")
	callStartCmd.Flags().BoolVar(&teams, "teams", false, "Teams call")
	callCmd.AddCommand(callStartCmd)
}
