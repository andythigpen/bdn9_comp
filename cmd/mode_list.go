package cmd

import (
	"fmt"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/spf13/cobra"
)

// modeListCmd represents the mode set command
var modeListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the RGB matrix animation modes",
	RunE: func(cmd *cobra.Command, args []string) error {
		var i pb.RGBMode
		for i = pb.RGBMode_RGB_MATRIX_INVALID + 1; i < pb.RGBMode_RGB_MATRIX_MAX; i++ {
			fmt.Println(i.Description())
		}
		return nil
	},
}

func init() {
	modeCmd.AddCommand(modeListCmd)
}
