package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/prompts"
	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook <name>",
	Short: "Create a new React 'Hook'",
	Args:  cobra.MinimumNArgs(1),
	Run:   hookCmdRun,
}

var hookOpt *models.CmdOptionsModel

func init() {
	hookOpt = models.NewCmdOptions()

	rootCmd.AddCommand(hookCmd)

	hookCmd.Flags().StringVarP(&hookOpt.Dest, "dest", "d", "", "destination directory")
	hookCmd.Flags().BoolP("interactive", "i", false, "use the simple interactive mode")
	hookCmd.Flags().BoolVarP(&hookOpt.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	hookCmd.Flags().BoolVarP(&hookOpt.WithTest, "with-test", "t", false, "create a unit test file")

	hookCmd.Flags().SortFlags = false
}

func hookCmdRun(cmd *cobra.Command, args []string) {
	hookOpt.Type = "hook"
	hookOpt.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); !interactive {
		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			hookOpt.Dest = "./hooks"
		}

		PrepareHookGeneration(hookOpt)
		return
	}

	prompts.NamePrompt(hookOpt)
	prompts.DestPrompt(hookOpt)
	prompts.UsesKebabCasePrompt(hookOpt)
	prompts.WithTestPrompt(hookOpt)

	PrepareHookGeneration(hookOpt)
}

func PrepareHookGeneration(opt *models.CmdOptionsModel) {
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
