package cmd

import (
	"context"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// echoCmd represents the reset command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Tests the keyboard connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.Echo(context.Background(), &pb.EchoRequest{})
		return err
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
