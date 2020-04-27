package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available go versions",
	RunE: func(cmd *cobra.Command, args []string) error {
		versions, err := availableVersions()
		if err != nil {
			return err
		}
		for i := len(versions)/2 - 1; i >= 0; i-- {
			opp := len(versions) - 1 - i
			versions[i], versions[opp] = versions[opp], versions[i]
		}
		for _, v := range versions {
			fmt.Println(v.String())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
