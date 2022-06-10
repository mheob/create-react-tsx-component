package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/component"
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

func init() {
	rootCmd.AddCommand(componentCmd)

	componentCmd.Flags().StringVarP(&models.CmdOptions.Dest, "dest", "d", "", "destination directory")
	componentCmd.Flags().BoolP("interactive", "i", false, "use the simple interactive mode")
	componentCmd.Flags().BoolVarP(&models.CmdOptions.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	componentCmd.Flags().BoolVarP(&models.CmdOptions.ShouldSkipStorybook, "no-storybook", "s", false, "don't create a storybook stories file")
	componentCmd.Flags().BoolVarP(&models.CmdOptions.ShouldSkipTest, "no-test", "t", false, "don't create a unit test file")

	componentCmd.Flags().SortFlags = false
}

func componentCmdRun(cmd *cobra.Command, args []string) {
	models.CmdOptions.Type = "component"
	models.CmdOptions.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); !interactive {
		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			models.CmdOptions.Dest = "./components"
		}

		component.Run()
		return
	}

	prompts.NamePrompt()
	prompts.DestPrompt()
	prompts.UsesKebabCasePrompt()
	prompts.ShouldSkipStorybookPrompt()
	prompts.ShouldSkipTestPrompt()

	component.Run()
}
