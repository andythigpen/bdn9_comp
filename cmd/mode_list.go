package cmd

import (
	"fmt"

	"github.com/andythigpen/bdn9_comp/v2/serial"
	"github.com/spf13/cobra"
)

// modeListCmd represents the mode set command
var modeListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the RGB matrix animation modes",
	RunE: func(cmd *cobra.Command, args []string) error {
		var i serial.RGBMode
		for i = serial.RGB_MATRIX_INVALID + 1; i < serial.RGB_MATRIX_MAX; i++ {
			fmt.Println(i.String())
		}
		return nil
	},
}

func init() {
	modeCmd.AddCommand(modeListCmd)
}
