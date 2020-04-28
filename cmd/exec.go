package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Use the selected go version",
	RunE: func(cmd *cobra.Command, args []string) error {
		goCMD := "go"
		version := viper.GetString("go-version")

		if version != "" {
			goCMD = "go" + version
		} else {
			cwd, err := os.Getwd()
			if err == nil {
				version, ok, err := getWorkspace(cwd)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
				if ok {
					goCMD = "go" + version

				}
			}
		}

		err := Exec(goCMD, args...)
		if err, ok := err.(*exec.ExitError); ok {
			os.Exit(err.ExitCode())
		}
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
