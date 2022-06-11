package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/prompts"
	"github.com/spf13/cobra"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component <name>",
	Short: "Create a new React 'Component'",
	Args:  cobra.MinimumNArgs(1),
	Run:   componentCmdRun,
}

var componentOpt *models.CmdOptionsModel

func init() {
	componentOpt = models.NewCmdOptions()

	rootCmd.AddCommand(componentCmd)

	componentCmd.Flags().BoolVarP(&componentOpt.WithStorybook, "with-storybook", "s", false, "create a storybook stories file")
	componentCmd.Flags().BoolVarP(&componentOpt.WithTest, "with-test", "t", false, "create a unit test file")

	componentCmd.Flags().SortFlags = false
}

func componentCmdRun(cmd *cobra.Command, args []string) {
	componentOpt.Type = "component"
	componentOpt.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); interactive {
		prompts.NamePrompt(componentOpt)
		prompts.DestPrompt(componentOpt)
		prompts.UsesKebabCasePrompt(componentOpt)
		prompts.WithStorybookPrompt(componentOpt)
		prompts.WithTestPrompt(componentOpt)

		PrepareComponentGeneration(componentOpt)
		return
	}

	if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
		componentOpt.Dest = "./components"
	}

	if usesKebabCase, _ := cmd.Flags().GetBool("kebab-case"); usesKebabCase {
		componentOpt.UsesKebabCase = usesKebabCase
	}

	PrepareComponentGeneration(componentOpt)
}

func PrepareComponentGeneration(opt *models.CmdOptionsModel) {
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
