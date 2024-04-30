package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

// hsvCmd represents the hsv command
var hsvCmd = &cobra.Command{
	Use:   "hsv",
	Short: "HSV color",
}

func parseHSV(args []string) (h, s, v uint64, err error) {
	if h, err = strconv.ParseUint(args[0], 10, 8); err != nil {
		return
	}
	if s, err = strconv.ParseUint(args[1], 10, 8); err != nil {
		return
	}
	if v, err = strconv.ParseUint(args[2], 10, 8); err != nil {
		return
	}
	return
}

func parseRGB(args []string) (r, g, b uint64, err error) {
	return parseHSV(args)
}

func init() {
	rootCmd.AddCommand(hsvCmd)
}
