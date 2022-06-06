package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/component"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/utils"
	"github.com/spf13/cobra"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use: "component <name>",
	Example: `  crtc component IconButton
  crtc component Icon Button -kst
  crtc component icon-button -k -s -t
  crtc component "Icon Button"
  crtc component icon button --dest ./components`,
	Short: "Create a new React 'Component'",
	Long:  "Create a new React 'Component'.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		models.CmdOptions.Name = strings.Join(args, " ")

		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			models.CmdOptions.Dest = utils.GetWd() + "/components"
		}

		component.Run()
	},
}

func init() {
	rootCmd.AddCommand(componentCmd)

	componentCmd.Flags().StringVarP(&models.CmdOptions.Dest, "dest", "d", "", "destination directory")
	componentCmd.Flags().BoolVarP(&models.CmdOptions.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	componentCmd.Flags().BoolVarP(&models.CmdOptions.WithoutStorybook, "no-storybook", "s", false, "don't create a storybook stories file")
	componentCmd.Flags().BoolVarP(&models.CmdOptions.WithoutTest, "no-test", "t", false, "don't create a unit test file")
	componentCmd.Flags().BoolVar(&models.CmdOptions.OnlyDryRun, "dry-run", false, "print only the changes, but don't write to disk")

	componentCmd.Flags().SortFlags = false
}
