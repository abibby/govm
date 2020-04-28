package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set the go version to use in the current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		ok, err := VersionExists("go" + args[0])
		if err != nil {
			return err
		}
		if !ok {
			fmt.Fprintf(os.Stderr, "\"%s\" is not a valid go version", args[0])
			os.Exit(1)
		}
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
