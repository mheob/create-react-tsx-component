package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/component"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/spf13/cobra"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Create a new React 'Component'",
	Long:  "Create a new React 'Component'.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dest, _ := cmd.Flags().GetString("dest")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		noStorybook, _ := cmd.Flags().GetBool("no-storybook")
		kebabCase, _ := cmd.Flags().GetBool("kebab-case")
		noTest, _ := cmd.Flags().GetBool("no-test")

		component.Init(&models.CmdOptions{
			Name:      strings.Join(args, " "),
			Dest:      dest,
			DryRun:    dryRun,
			KebabCase: kebabCase,
			Storybook: !noStorybook,
			Test:      !noTest,
		})
	},
}

func init() {
	rootCmd.AddCommand(componentCmd)
	componentCmd.Flags().StringP("dest", "d", "", "destination directory")
	componentCmd.Flags().Bool("dry-run", false, "Print only the changes, but don't write to disk")
	componentCmd.Flags().BoolP("kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	componentCmd.Flags().BoolP("no-storybook", "s", false, "don't create a storybook stories file")
	componentCmd.Flags().BoolP("no-test", "t", false, "don't create a unit test file")
}
