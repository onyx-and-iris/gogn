// Package main provides the root command for the gogn CLI application.
package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/charmbracelet/fang"
	"github.com/cuonglm/gogi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string // Version of the CLI, set during build time

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "gogn",
	Short: "A command-line interface for generating .gitignore files",
	Long: `gogn is a command-line interface for generating .gitignore files.
It allows users to easily create and manage .gitignore files for various programming languages and frameworks.
You may also list available templates and generate .gitignore files based on those templates.`,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		client, err := gogi.NewHTTPClient()
		if err != nil {
			return fmt.Errorf("error creating HTTP client: %w", err)
		}
		cmd.SetContext(WithClient(cmd.Context(), client))
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().IntP("height", "H", 10, "Height of the selection prompt")
	rootCmd.PersistentFlags().
		StringP("filter", "f", "startswith", "Type of filter to apply to the list of templates (e.g., 'startswith', 'contains')")
	rootCmd.PersistentFlags().BoolP("start-search", "s", false, "Start the prompt in search mode")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("GOGN")
	viper.AutomaticEnv()
	viper.BindPFlag("height", rootCmd.PersistentFlags().Lookup("height"))
	viper.BindPFlag("filter", rootCmd.PersistentFlags().Lookup("filter"))
	viper.BindPFlag("start-search", rootCmd.PersistentFlags().Lookup("start-search"))
}

// main is the entry point of the application.
// It executes the root command and handles any errors.
func main() {
	if err := fang.Execute(context.Background(), rootCmd, fang.WithVersion(versionFromBuild())); err != nil {
		os.Exit(1)
	}
}

func versionFromBuild() string {
	if version == "" {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			return "(unable to read version)"
		}
		version = strings.Split(info.Main.Version, "-")[0]
	}

	return version
}
