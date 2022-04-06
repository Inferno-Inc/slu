package df

import (
	"fmt"
	"log"
	"strings"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "df",
	Short: "System's \"df\" filtered for /dev devices and human readable",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		out, err := exec_utils.ExecStr("df", "-h")
		if err != nil {
			log.Fatalln(err)
		}
		lines := strings.Split(out, "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, "/dev") || i == 0 {
				fmt.Println(line)
			}
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
