package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/page"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/utils"
	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use: "page <name>",
	Example: `  crtc page MyPage
  crtc page MyPage -kt
  crtc page My Page -k -t
  crtc page "My Page"
  crtc page my page --dest ./pages`,
	Short: "Create a new React 'Page'",
	Long:  "Create a new React 'Page'.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		models.CmdOptions.Name = strings.Join(args, " ")

		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			models.CmdOptions.Dest = utils.GetWd() + "/pages"
		}

		page.Run()
	},
}

func init() {
	rootCmd.AddCommand(pageCmd)

	pageCmd.Flags().StringVarP(&models.CmdOptions.Dest, "dest", "d", "", "destination directory")
	pageCmd.Flags().BoolVarP(&models.CmdOptions.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	pageCmd.Flags().BoolVar(&models.CmdOptions.OnlyDryRun, "dry-run", false, "print only the changes, but don't write to disk")

	pageCmd.Flags().SortFlags = false
}
