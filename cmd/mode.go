package cmd

import (
	"github.com/spf13/cobra"
)

// modeCmd represents the mode command
var modeCmd = &cobra.Command{
	Use:   "mode",
	Short: "RGB matrix animation modes",
}

func init() {
	rootCmd.AddCommand(modeCmd)
}
