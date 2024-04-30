package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/andythigpen/bdn9_comp/v2/device"
	pb "github.com/andythigpen/bdn9_comp/v2/proto"
)

var (
	cfgFile  string
	keyboard device.BDN9Device
	conn     *grpc.ClientConn
	client   pb.BDN9ServiceClient
	persist  bool
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

	if keyboard != nil && keyboard.IsOpen() {
		keyboard.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func init() {
	cobra.OnInitialize(InitConfig, OpenDevice)

	path, _ := xdg.ConfigFile("bdn9")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default: %s/config.yaml)", path))
	rootCmd.PersistentFlags().BoolVarP(&persist, "persist", "p", false, "Write changes to EEPROM")
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
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

	// set defaults
	viper.SetDefault("bind", "localhost:17432")
	viper.SetDefault("slackWindowName", "slack")
	viper.SetDefault("slackWindowTitle", "slack")
	viper.SetDefault("slackMuteKeys", []string{"space", "lcmd", "lshift"})
	viper.SetDefault("teamsWindowName", "teams")
	viper.SetDefault("teamsWindowTitle", "microsoft teams")
	viper.SetDefault("teamsMuteKeys", []string{"m", "lcmd", "lshift"})
}

func OpenDevice() {
	server := viper.GetString("server")
	if len(server) != 0 {
		openGrpc()
	} else {
		if _, err := OpenUSBDevice(nil); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
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

func OpenUSBDevice(handler device.EventHandler) (device.BDN9Device, error) {
	keyboard = device.NewDevice(handler)
	if err := keyboard.Open(); err != nil {
		return nil, err
	}

	if persist {
		keyboard.EnablePersist()
	}

	client = device.NewDeviceClient(keyboard)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if err := keyboard.Close(); err != nil {
			fmt.Printf("failed to close keyboard connection: %s", err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
	return keyboard, nil
}
