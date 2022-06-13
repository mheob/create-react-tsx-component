package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/models"
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

	hookCmd.Flags().BoolVarP(&hookOpt.WithTest, "with-test", "t", false, "create a unit test file")

	hookCmd.Flags().SortFlags = false
}

func hookCmdRun(cmd *cobra.Command, args []string) {
	hookOpt.Type = "hook"
	hookOpt.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); interactive {
		models.NamePrompt(hookOpt)
		models.DestPrompt(hookOpt)
		models.UsesKebabCasePrompt(hookOpt)
		models.WithTestPrompt(hookOpt)

		PrepareHookGeneration(hookOpt)
		return
	}

	if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
		hookOpt.Dest = "./src/hooks"
	}

	if usesKebabCase, _ := cmd.Flags().GetBool("kebab-case"); usesKebabCase {
		hookOpt.UsesKebabCase = usesKebabCase
	}

	PrepareHookGeneration(hookOpt)
}

func PrepareHookGeneration(opt *models.CmdOptionsModel) {
	opt.SetNames()

	var vars = make(models.TmplVars)
	vars["Name"] = opt.ReactName
	vars["FileName"] = opt.FileName

	g := models.NewGenerator("hook", opt, vars)

	files := make([]models.File, 1, 2)
	files[0] = models.File{Name: "hook", Extension: ".tsx"}

	if opt.WithTest {
		files = append(files, models.File{Name: "test", Extension: ".test.ts"})
	}

	g.GenerateFiles(files)
}
