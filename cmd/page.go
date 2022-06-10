package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/page"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/prompts"
	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page <name>",
	Short: "Create a new React 'Page'",
	Args:  cobra.MinimumNArgs(1),
	Run:   pageCmdRun,
}

func init() {
	rootCmd.AddCommand(pageCmd)

	pageCmd.Flags().StringVarP(&models.CmdOptions.Dest, "dest", "d", "", "destination directory")
	pageCmd.Flags().BoolP("interactive", "i", false, "use the simple interactive mode")
	pageCmd.Flags().BoolVarP(&models.CmdOptions.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")

	pageCmd.Flags().SortFlags = false
}

func pageCmdRun(cmd *cobra.Command, args []string) {
	models.CmdOptions.Type = "page"
	models.CmdOptions.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); !interactive {
		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			models.CmdOptions.Dest = "./pages"
		}

		page.Run()
		return
	}

	prompts.NamePrompt()
	prompts.DestPrompt()
	prompts.UsesKebabCasePrompt()

	page.Run()
}
