package cmd

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

// debugCmd helps finding window names
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Lists windows",
	RunE: func(cmd *cobra.Command, args []string) error {
		names, err := robotgo.FindNames()
		if err != nil {
			return err
		}
		for _, name := range names {
			fmt.Println(name)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
