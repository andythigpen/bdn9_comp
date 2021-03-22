package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/andythigpen/bdn9_comp/v2/serial"
)

var (
	cfgFile string
	device  serial.BDN9SerialDevice
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
}

func init() {
	cobra.OnInitialize(initConfig, openDevice)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bdn9.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&persist, "persist", "p", false, "Write changes to EEPROM")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bdn9" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bdn9")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func openDevice() {
	device = serial.NewDevice()
	if err := device.Open(viper.GetString("port")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if persist {
		device.EnablePersist()
	}
}
