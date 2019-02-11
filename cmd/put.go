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
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	resty "gopkg.in/resty.v1"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		somebody := new(Somebody)
		j := strings.Join(args, ",")
		s := strings.Split(j, ",")
		fmt.Println("copy args to new slice:", s, "len:", len(s))
		if len(s) < 3 {
			fmt.Errorf("Invalir argument! args format is questionID,answerID,nickname")
		}
		somebody.Name = s[0]
		somebody.Email = s[1]
		b, _ := strconv.ParseBool(s[2])
		somebody.Active = b
		somebody.Balance = s[3]
		somebody.Picture = s[4]
		somebody.Age = s[5]
		somebody.EyeColor = s[6]
		somebody.Gender = s[7]
		somebody.Company = s[8]
		somebody.Phone = s[9]
		somebody.Address = s[10]
		somebody.About = s[11]

		body, _ := json.Marshal(somebody)

		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8080/demoapi/1/1/1/signup/appName/")

		log.Println("Error:", err)
		log.Println("Response Status Code:", resp.StatusCode())
		log.Println("Response Body:", resp)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
