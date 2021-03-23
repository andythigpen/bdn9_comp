package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// callEndCmd represents the call end command
var callEndCmd = &cobra.Command{
	Use:   "end",
	Short: "Ends call mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.EndCall(context.Background(), &pb.EndCallRequest{})
		return err
	},
}

func init() {
	callCmd.AddCommand(callEndCmd)
}
