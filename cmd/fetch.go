// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
			RootCmd.Help()
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
