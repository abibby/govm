package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set the go version to use in the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		return setWorkspace(cwd, args[0])
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
