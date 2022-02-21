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
		fmt.Println("Slack:")
		pids, err := robotgo.FindIds("slack")
		for _, pid := range pids {
			fmt.Printf("%d %s\n", pid, robotgo.GetTitle(pid))
		}
		fmt.Println("")

		fmt.Println("Teams:")
		pids, err = robotgo.FindIds("teams")
		for _, pid := range pids {
			fmt.Printf("%d %s\n", pid, robotgo.GetTitle(pid))
		}
		fmt.Println("")

		names, err := robotgo.FindNames()
		if err != nil {
			return err
		}
		fmt.Println("All names:")
		for _, name := range names {
			fmt.Println(name)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
