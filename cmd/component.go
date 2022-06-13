package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/models"
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
		models.NamePrompt(componentOpt)
		models.DestPrompt(componentOpt)
		models.UsesKebabCasePrompt(componentOpt)
		models.WithStorybookPrompt(componentOpt)
		models.WithTestPrompt(componentOpt)

		PrepareComponentGeneration(componentOpt)
		return
	}

	if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
		componentOpt.Dest = "./src/components"
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

	files := make([]models.File, 1, 3)
	files[0] = models.File{Name: "component", Extension: ".tsx"}

	if opt.WithStorybook {
		files = append(files, models.File{Name: "stories", Extension: ".stories.tsx"})
	}

	if opt.WithTest {
		files = append(files, models.File{Name: "test", Extension: ".test.tsx"})
	}

	g.GenerateFiles(files)
}
