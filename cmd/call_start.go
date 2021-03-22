package cmd

import (
	"github.com/andythigpen/bdn9_comp/v2/serial"
	"github.com/spf13/cobra"
)

var slack, teams bool

// callStartCmd represents the call start command
var callStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a call mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		if teams {
			if err := device.ActivateLayer(serial.LAYER_TEAMS); err != nil {
				return err
			}
			return device.SetMuteStatus(true) // teams starts muted
		} else {
			if err := device.ActivateLayer(serial.LAYER_SLACK); err != nil {
				return err
			}
			return device.SetMuteStatus(false) // slack starts unmuted
		}
	},
}

func init() {
	callStartCmd.Flags().BoolVar(&slack, "slack", false, "Slack call")
	callStartCmd.Flags().BoolVar(&teams, "teams", false, "Teams call")
	callCmd.AddCommand(callStartCmd)
}
