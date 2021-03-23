package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the keyboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.Reset(context.Background(), &pb.ResetRequest{})
		return err
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
