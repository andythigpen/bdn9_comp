package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// modeToggleCmd represents the mode toggle command
var modeToggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggles the RGB matrix",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.ToggleMatrix(context.Background(), &pb.ToggleMatrixRequest{})
		return err
	},
}

func init() {
	modeCmd.AddCommand(modeToggleCmd)
}
