// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
	resty "gopkg.in/resty.v1"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		j := strings.Join(args, ",")
		s := strings.Split(j, ",")
		fmt.Println("copy args to new slice:", s, "len:", len(s))
		if len(s) < 1 {
			fmt.Errorf("Invalir argument! args format is questionID,answerID,nickname")
		}
		var query string
		dataInput := s[0]
		if strings.Contains(dataInput, "@") {
			query = "?email=" + dataInput
		} else {
			query = "?id=" + dataInput
		}
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Delete("http://localhost:8080/demoapi/1/1/1/delete/appName/" + query)

		log.Println("Error:", err)
		log.Println("Response Status Code:", resp.StatusCode(), " query:", query)
		log.Println("Response Body:", resp)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
