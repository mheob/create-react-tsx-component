package cmd

import (
	"fmt"
	"os"

	"github.com/mheob/create-react-tsx-component/generators/component"
	"github.com/mheob/create-react-tsx-component/generators/hook"
	"github.com/mheob/create-react-tsx-component/generators/page"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/prompts"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "crtc",
	Version:           "1.0.0",
	Short:             "Use 'crtc' to create a new React Component",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Run:               rootCmdRun,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// disable default help command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.PersistentFlags().BoolVar(&models.CmdOptions.OnlyDryRun, "dry-run", false, "print only the changes, but don't write to disk")
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	prompts.TypePrompt()
	prompts.NamePrompt()
	prompts.DestPrompt()
	prompts.UsesKebabCasePrompt()

	switch models.CmdOptions.Type {
	case "component":
		prompts.ShouldSkipStorybookPrompt()
		prompts.ShouldSkipTestPrompt()
		component.Run()
	case "hook":
		prompts.ShouldSkipTestPrompt()
		hook.Run()
	case "page":
		page.Run()
	default:
		fmt.Println("Invalid choice")
	}
}
