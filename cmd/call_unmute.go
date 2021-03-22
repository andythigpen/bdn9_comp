package cmd

import (
	"github.com/spf13/cobra"
)

// callUnmuteCmd represents the call unmute command
var callUnmuteCmd = &cobra.Command{
	Use:   "unmute",
	Short: "Unmutes call mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return device.SetMuteStatus(false)
	},
}

func init() {
	callCmd.AddCommand(callUnmuteCmd)
}
