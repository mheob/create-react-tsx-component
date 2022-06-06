package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "crtc",
	Version: "1.0.0",
	Short:   "Use 'crtc' to create a new React Component",
	Long: `Use 'crtc' to create a new React Component.
You can also add new React Pages, Hooks and more.`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// disable default help command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
