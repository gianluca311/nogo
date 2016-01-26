package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/gianluca311/nogo/notam"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch [ICAO-Codes]",
	Short: "fetches the NOTAMs",
	Long: `Fetches the NOTAMs of the given ICAO-Code
    from the source "FAA" and returns it to you.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
		}
		for _, arg := range args {
			if len(arg) != 4 {
				log.Fatalln(arg, "is not a valid ICAO code")
			}
			arg := strings.ToUpper(arg)
			log.Println("fetching now for", arg)
			notamList := notam.GetNotams(arg)
			for _, elem := range notamList.Notamlist {
				fmt.Println(elem)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)
}
