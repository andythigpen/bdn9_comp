package cmd

import (
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/andythigpen/bdn9_comp/v2/service"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs an RPC server",
	RunE: func(cmd *cobra.Command, args []string) error {
		bindAddress := viper.GetString("bind")
		if len(bindAddress) == 0 {
			bindAddress = "localhost:17432"
		}
		lis, err := net.Listen("tcp", bindAddress)
		if err != nil {
			return err
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		server := service.NewService(device)
		pb.RegisterBDN9ServiceServer(grpcServer, server)
		return grpcServer.Serve(lis)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
