package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/hook"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/prompts"
	"github.com/mheob/create-react-tsx-component/utils"
	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook <name>",
	Short: "Create a new React 'Hook'",
	Args:  cobra.MinimumNArgs(1),
	Run:   hookCmdRun,
}

func init() {
	rootCmd.AddCommand(hookCmd)

	hookCmd.Flags().StringVarP(&models.CmdOptions.Dest, "dest", "d", "", "destination directory")
	hookCmd.Flags().BoolP("interactive", "i", false, "use the simple interactive mode")
	hookCmd.Flags().BoolVarP(&models.CmdOptions.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	hookCmd.Flags().BoolVarP(&models.CmdOptions.ShouldSkipTest, "no-test", "t", false, "don't create a unit test file")
	hookCmd.Flags().BoolVar(&models.CmdOptions.OnlyDryRun, "dry-run", false, "print only the changes, but don't write to disk")

	hookCmd.Flags().SortFlags = false
}

func hookCmdRun(cmd *cobra.Command, args []string) {
	models.CmdOptions.Type = "hook"
	models.CmdOptions.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); interactive != true {
		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			models.CmdOptions.Dest = utils.GetWd() + "/hooks"
		}

		hook.Run()
		return
	}

	prompts.NamePrompt()
	prompts.DestPrompt()
	prompts.UsesKebabCasePrompt()
	prompts.ShouldSkipTestPrompt()

	hook.Run()
}
