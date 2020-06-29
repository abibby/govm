package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "govm",
	Short: "Manage and use multiple versions of go",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().StringP("go-version", "g", "", "config file")
	viper.BindPFlag("go-version", rootCmd.PersistentFlags().Lookup("go-version"))

	// viper.SetDefault("go-version", )

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		cfgDir, err := os.UserConfigDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetDefault("workspaces", path.Join(cfgDir, "govm", "workspaces.json"))
		// Search config in home directory with name ".govm" (without extension).
		viper.AddConfigPath(cfgDir)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.ReadInConfig()
}

func readWorkspaces() (map[string]string, error) {
	workspaces := viper.GetString("workspaces")

	if _, err := os.Stat(workspaces); os.IsNotExist(err) {
		return map[string]string{}, nil
	}

	b, err := ioutil.ReadFile(workspaces)
	if err != nil {
		return nil, err
	}

	versions := map[string]string{}

	err = json.Unmarshal(b, &versions)
	if err != nil {
		return nil, err
	}
	return versions, nil
}
func writeWorkspaces(versions map[string]string) error {
	b, err := json.MarshalIndent(versions, "", "    ")
	if err != nil {
		return err
	}
	workspaces := viper.GetString("workspaces")
	dir, _ := filepath.Split(workspaces)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(workspaces, b, 0644)
}

func getWorkspace(cwd string) (string, bool, error) {
	versions, err := readWorkspaces()
	if err != nil {
		return "", false, err
	}

	currentWorkspaceDir := ""
	currentVersion := ""

	for workspaceDir, version := range versions {
		if isInDir(workspaceDir, cwd) {
			if len(workspaceDir) > len(currentWorkspaceDir) {
				currentWorkspaceDir = workspaceDir
				currentVersion = version
			}
		}
	}

	if currentVersion != "" {
		return currentVersion, true, nil
	}

	return "", false, nil
}

func setWorkspace(workspaceDir, version string) error {
	versions, err := readWorkspaces()
	if err != nil {
		return err
	}
	versions[workspaceDir] = version
	return writeWorkspaces(versions)
}

func isInDir(dir, file string) bool {
	dirParts := strings.Split(dir, string(os.PathSeparator))
	fileParts := strings.Split(file, string(os.PathSeparator))

	if len(dirParts) > len(fileParts) {
		return false
	}

	for i, part := range dirParts {
		if part != fileParts[i] {
			return false
		}
	}
	return true
}

func Exec(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
