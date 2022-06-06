package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/page"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Create a new React 'Page'",
	Long:  "Create a new React 'Page'.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dest, _ := cmd.Flags().GetString("dest")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		kebabCase, _ := cmd.Flags().GetBool("kebab-case")

		page.Init(&models.CmdOptions{
			Name:      strings.Join(args, " "),
			Dest:      dest,
			DryRun:    dryRun,
			KebabCase: kebabCase,
		})
	},
}

func init() {
	rootCmd.AddCommand(pageCmd)
	pageCmd.Flags().Bool("dry-run", false, "print only the changes, but don't write to disk")
	pageCmd.Flags().BoolP("kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
}
