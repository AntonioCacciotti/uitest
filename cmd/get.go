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
	"strings"

	"github.com/spf13/cobra"
	resty "gopkg.in/resty.v1"
)

//company=RODEOLOGY,eyeColor=blue,age=35&limit=1000
// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get some data from the database",
	Run: func(cmd *cobra.Command, args []string) {
		joinArgs := strings.Join(args, "")
		s := strings.Split(joinArgs, ",")
		query := s[0]
		limit := s[2]
		if query == "company" {
			query = "?company=" + s[1]
		} else if query == "eyeColor" {
			query = "?eyeColor=" + s[1]
		} else if query == "age" {
			query = "?age=" + s[1]
		} else {
			limit = "10"
		}
		query = query + "&limit=" + limit
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get("http://localhost:8080/demoapi/1/1/1/getsomebody" + query)

		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode(), "query:", query)
		fmt.Printf("\nResponse Body: %v", resp)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
