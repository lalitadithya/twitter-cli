/*
Copyright © 2020 Lalit Adithya <lalit@lalitadithya.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/lalitadithya/twitter-cli/twitter"
	"github.com/lalitadithya/twitter-cli/twitter/util"
	"github.com/spf13/cobra"
)

// authorizeCmd represents the authorize command
var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("authorize called")
		processAuthorize()
	},
}

func processAuthorize() {
	client, err := twitter.LoadSecrets()
	if err != nil {
		if errors.Is(err, util.APIKeyMissingError) {
			client = setClientSecrets()
		} else {
			panic(err)
		}
	}

	fmt.Println(client)
}

func setClientSecrets() *twitter.TwitterClient {
	fmt.Println("Need to get API keys")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter API Key: ")
	apiKey, _ := reader.ReadString('\n')

	fmt.Print("Enter API secret Key: ")
	apiSecretKey, _ := reader.ReadString('\n')

	client, err := twitter.SetSecrets(apiKey, apiSecretKey)
	if err != nil {
		fmt.Println("Unable to set secrets ", err)
		os.Exit(1)
	}

	return client
}

func init() {
	rootCmd.AddCommand(authorizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authorizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authorizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
