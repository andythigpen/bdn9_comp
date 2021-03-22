package cmd

import (
	"github.com/spf13/cobra"
)

// layerCmd represents the layer command
var layerCmd = &cobra.Command{
	Use:   "layer",
	Short: "Active keyboard layer(s)",
}

func init() {
	rootCmd.AddCommand(layerCmd)
}
