package cmd

import (
	"github.com/spf13/cobra"
)

// modeToggleCmd represents the mode toggle command
var modeToggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggles the RGB matrix",
	RunE: func(cmd *cobra.Command, args []string) error {
		return device.ToggleMatrix()
	},
}

func init() {
	modeCmd.AddCommand(modeToggleCmd)
}
