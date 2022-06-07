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
		break
	case "hook":
		prompts.ShouldSkipTestPrompt()
		hook.Run()
		break
	case "page":
		page.Run()
		break
	default:
		fmt.Println("Invalid choice")
	}
}
