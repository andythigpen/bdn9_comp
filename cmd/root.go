package cmd

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/andythigpen/bdn9_comp/v2/serial"
)

var (
	cfgFile string
	device  serial.BDN9SerialDevice
	conn    *grpc.ClientConn
	client  pb.BDN9ServiceClient
	persist bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bdn9",
	Short: "CLI for BDN9",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if device != nil && device.IsOpen() {
		device.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func init() {
	cobra.OnInitialize(initConfig, openDevice)

	path, _ := xdg.ConfigFile("bdn9")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default: %s/config.yaml)", path))
	rootCmd.PersistentFlags().BoolVarP(&persist, "persist", "p", false, "Write changes to EEPROM")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		path, err := xdg.ConfigFile("bdn9")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(path)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func openDevice() {
	port := viper.GetString("port")
	if len(port) == 0 {
		openGrpc()
	} else {
		openSerialDevice()
	}
}

func openGrpc() {
	var err error
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err = grpc.Dial(viper.GetString("server"), opts...)
	if err != nil {
		panic(err)
	}
	client = pb.NewBDN9ServiceClient(conn)
}

func openSerialDevice() {
	port := viper.GetString("port")
	device = serial.NewDevice()
	if err := device.Open(port); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if persist {
		device.EnablePersist()
	}

	client = serial.NewSerialClient(device)
}
