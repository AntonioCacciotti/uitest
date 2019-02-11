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

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Create a new record using sintex [name,email,balance,picture,age,eyeColor,gender,company,address]",
	Run: func(cmd *cobra.Command, args []string) {
		somebody := new(Somebody)
		j := strings.Join(args, ",")
		s := strings.Split(j, ",")
		fmt.Println("copy args to new slice:", s, "len:", len(s))
		if len(s) < 3 {
			fmt.Errorf("Invalid argument! args format is questionID,answerID,nickname")
		}
		somebody.Name = LenghtArray(s, 0)
		somebody.Email = LenghtArray(s, 1)
		somebody.Active = true
		somebody.Balance = LenghtArray(s, 2)
		somebody.Picture = LenghtArray(s, 3)
		somebody.Age = LenghtArray(s, 4)
		somebody.EyeColor = LenghtArray(s, 5)
		somebody.Gender = LenghtArray(s, 6)
		somebody.Company = LenghtArray(s, 7)
		somebody.Phone = LenghtArray(s, 8)
		somebody.Address = LenghtArray(s, 9)
		somebody.About = LenghtArray(s, 10)

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

//LenghtArray get the element in the b position
func LenghtArray(a []string, b int) string {
	if len(a) > b {
		return a[b]
	}
	return ""
}
func init() {
	rootCmd.AddCommand(postCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func convertStringToIng(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("convertsion string to int FATAL ERROR")
	}
	return i
}

//Somebody obj
type Somebody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Active   bool   `json:"isActive"`
	Balance  string `json:"balance"`
	Picture  string `json:"picture"`
	Age      string `json:"age"`
	EyeColor string `json:"eyeColor"`
	Gender   string `json:"gender"`
	Company  string `json:"company"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	About    string `json:"about"`
}
