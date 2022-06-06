package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/hook"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Create a new React 'Hook'",
	Long:  "Create a new React 'Hook'.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dest, _ := cmd.Flags().GetString("dest")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		kebabCase, _ := cmd.Flags().GetBool("kebab-case")
		noTest, _ := cmd.Flags().GetBool("no-test")

		hook.Init(&models.CmdOptions{
			Name:      strings.Join(args, " "),
			Dest:      dest,
			DryRun:    dryRun,
			KebabCase: kebabCase,
			Test:      !noTest,
		})
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
	hookCmd.Flags().Bool("dry-run", false, "Print only the changes, but don't write to disk")
	hookCmd.Flags().BoolP("kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	hookCmd.Flags().BoolP("no-test", "t", false, "don't create a unit test file")
}
