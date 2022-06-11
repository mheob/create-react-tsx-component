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

var rootOpt *models.CmdOptionsModel

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootOpt = models.NewCmdOptions()

	// disable default help command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.PersistentFlags().BoolVar(&rootOpt.OnlyDryRun, "dry-run", false, "print only the changes, but don't write to disk")
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	prompts.TypePrompt(rootOpt)
	prompts.NamePrompt(rootOpt)
	prompts.DestPrompt(rootOpt)
	prompts.UsesKebabCasePrompt(rootOpt)

	switch rootOpt.Type {
	case "component":
		prompts.WithStorybookPrompt(rootOpt)
		prompts.WithTestPrompt(rootOpt)
		component.Run(rootOpt)
	case "hook":
		prompts.WithTestPrompt(rootOpt)
		hook.Run(rootOpt)
	case "page":
		page.Run(rootOpt)
	default:
		fmt.Println("Invalid choice")
	}
}
