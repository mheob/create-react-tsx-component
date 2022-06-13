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
		pageOpt.Dest = "./pages"
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
	vars["FileName"] = opt.FileName
	vars["WithTest"] = opt.WithTest

	g := models.NewGenerator("component", opt, vars)

	files := make([]string, 1, 3)
	files[0] = "component"

	if opt.WithStorybook {
		files = append(files, "stories")
	}

	if opt.WithTest {
		files = append(files, "test")
	}

	g.GenerateFiles(files)
}
