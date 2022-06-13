package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/models"
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

	pageCmd.Flags().SortFlags = false
}

func pageCmdRun(cmd *cobra.Command, args []string) {
	pageOpt.Type = "page"
	pageOpt.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); interactive {
		models.NamePrompt(pageOpt)
		models.DestPrompt(pageOpt)
		models.UsesKebabCasePrompt(pageOpt)

		PreparePageGeneration(pageOpt)
		return
	}

	if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
		pageOpt.Dest = "./src/pages"
	}

	if usesKebabCase, _ := cmd.Flags().GetBool("kebab-case"); usesKebabCase {
		pageOpt.UsesKebabCase = usesKebabCase
	}

	PreparePageGeneration(pageOpt)
}

func PreparePageGeneration(opt *models.CmdOptionsModel) {
	opt.SetNames()

	var vars = make(models.TmplVars)
	vars["Name"] = opt.ReactName

	g := models.NewGenerator("page", opt, vars)

	files := make([]models.File, 1)
	files[0] = models.File{Name: "page", Extension: ".tsx"}

	g.GenerateFiles(files)
}
