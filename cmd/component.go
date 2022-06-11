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

var componentOpt *models.CmdOptionsModel

func init() {
	componentOpt = models.NewCmdOptions()

	rootCmd.AddCommand(componentCmd)

	componentCmd.Flags().StringVarP(&componentOpt.Dest, "dest", "d", "", "destination directory")
	componentCmd.Flags().BoolP("interactive", "i", false, "use the simple interactive mode")
	componentCmd.Flags().BoolVarP(&componentOpt.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	componentCmd.Flags().BoolVarP(&componentOpt.WithStorybook, "with-storybook", "s", false, "create a storybook stories file")
	componentCmd.Flags().BoolVarP(&componentOpt.WithTest, "with-test", "t", false, "create a unit test file")

	componentCmd.Flags().SortFlags = false
}

func componentCmdRun(cmd *cobra.Command, args []string) {
	componentOpt.Type = "component"
	componentOpt.Name = strings.Join(args, " ")

	if interactive, _ := cmd.Flags().GetBool("interactive"); !interactive {
		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			componentOpt.Dest = "./components"
		}

		component.Run(componentOpt)
		return
	}

	prompts.NamePrompt(componentOpt)
	prompts.DestPrompt(componentOpt)
	prompts.UsesKebabCasePrompt(componentOpt)
	prompts.WithStorybookPrompt(componentOpt)
	prompts.WithTestPrompt(componentOpt)

	component.Run(componentOpt)
}
