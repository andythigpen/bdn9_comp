package cmd

import (
	"github.com/spf13/cobra"
)

// callEndCmd represents the call end command
var callEndCmd = &cobra.Command{
	Use:   "end",
	Short: "Ends call mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return device.EndCall()
	},
}

func init() {
	callCmd.AddCommand(callEndCmd)
}
