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

	"github.com/spf13/cobra"
)

// completeVersionsCmd represents the completeVersions command
var completeVersionsCmd = &cobra.Command{
	Use:   "completeVersions",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("1: :(")
		versions, err := availableVersions()
		if err == nil {
			for _, version := range versions {
				fmt.Printf(`"%s" `, version.String()[2:])
			}
		}
		fmt.Print(")")
	},
}

func init() {
	rootCmd.AddCommand(completeVersionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeVersionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeVersionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
