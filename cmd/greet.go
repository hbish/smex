/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Prints a greet message",
	Long: `Prints a greet message`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name")
		if name.Value.String() == "" {
			fmt.Printf("Hello World! %s\n", strings.Join(args, " "))
		} else {
			fmt.Printf("Hello %s, %s\n", name.Value.String(), strings.Join(args, " "))
		}
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
	greetCmd.Flags().StringP("name", "n", "", "The name to use")
}
