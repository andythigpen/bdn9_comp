package cmd

import (
	"github.com/spf13/cobra"
)

// callMuteCmd represents the call mute command
var callMuteCmd = &cobra.Command{
	Use:   "mute",
	Short: "Mutes call mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return device.SetMuteStatus(true)
	},
}

func init() {
	callCmd.AddCommand(callMuteCmd)
}
