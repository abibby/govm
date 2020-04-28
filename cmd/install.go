package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install <version>",
	Short: "Install a new version of go",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return installVersion(args[0])
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func installVersion(v string) error {
	goCMD := fmt.Sprintf("go%s", v)

	err := Exec("go", "get", fmt.Sprintf("golang.org/dl/%s", goCMD))
	if err != nil {
		return err
	}
	err = Exec(goCMD, "download")
	if err != nil {
		return err
	}
	return nil
}
