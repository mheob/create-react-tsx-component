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

var pageOpt *models.CmdOptionsModel

func init() {
	pageOpt = models.NewCmdOptions()

	rootCmd.AddCommand(pageCmd)

	pageCmd.Flags().StringVarP(&pageOpt.Dest, "dest", "d", "", "destination directory")
	pageCmd.Flags().BoolP("interactive", "i", false, "use the simple interactive mode")
	pageCmd.Flags().BoolVarP(&pageOpt.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")

	pageCmd.Flags().SortFlags = false
}

func pageCmdRun(cmd *cobra.Command, args []string) {
	pageOpt.Type = "page"
	pageOpt.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); !interactive {
		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			pageOpt.Dest = "./pages"
		}

		page.Run(pageOpt)
		return
	}

	prompts.NamePrompt(pageOpt)
	prompts.DestPrompt(pageOpt)
	prompts.UsesKebabCasePrompt(pageOpt)

	page.Run(pageOpt)
}
