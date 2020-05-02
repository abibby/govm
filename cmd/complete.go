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
	"os"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.GenZshCompletion(os.Stdout)
		fmt.Println("compdef _govm govm")
		fmt.Println(`function _govm_install {
	_arguments \
		'--config[config file]:' \
		'(-g --go-version)'{-g,--go-version}'[config file]:' \
		"$(govm completeVersions)"
}`)
		fmt.Println(`function _govm_use {
_arguments \
'--config[config file]:' \
'(-g --go-version)'{-g,--go-version}'[config file]:' \
"$(govm completeVersions)"
}`)
		fmt.Println("function go() { govm exec -- $@ }")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
