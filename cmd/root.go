package cmd

import (
	"fmt"
	"os"

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
	rootCmd.PersistentFlags().StringVarP(&rootOpt.Dest, "dest", "d", "", "destination directory")
	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "use the simple interactive mode")
	rootCmd.PersistentFlags().BoolVarP(&rootOpt.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	rootCmd.PersistentFlags().BoolVar(&rootOpt.OnlyDryRun, "dry-run", false, "print only the changes, but don't write to disk")

	rootCmd.PersistentFlags().SortFlags = false
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
		PrepareComponentGeneration(rootOpt)
	case "hook":
		prompts.WithTestPrompt(rootOpt)
		PrepareHookGeneration(rootOpt)
	case "page":
		PreparePageGeneration(rootOpt)
	default:
		fmt.Println("Invalid choice")
	}
}
