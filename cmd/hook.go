package cmd

import (
	"strings"

	"github.com/mheob/create-react-tsx-component/generators/hook"
	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/utils"
	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use: "hook <name>",
	Example: `  crtc hook My Hook
  crtc hook my-hook -kt
  crtc hook my Hook -k -t
  crtc hook "my hook"
  crtc hook my hook --dest ./hooks`,
	Short: "Create a new React 'Hook'",
	Long:  "Create a new React 'Hook'.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		models.CmdOptions.Name = strings.Join(args, " ")

		if dest, _ := cmd.Flags().GetString("dest"); dest == "" {
			models.CmdOptions.Dest = utils.GetWd() + "/hooks"
		}

		hook.Run()
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)

	hookCmd.Flags().StringVarP(&models.CmdOptions.Dest, "dest", "d", "", "destination directory")
	hookCmd.Flags().BoolVarP(&models.CmdOptions.UsesKebabCase, "kebab-case", "k", false, "use kebab-case instead of PascalCase for the filename")
	hookCmd.Flags().BoolVarP(&models.CmdOptions.WithoutTest, "no-test", "t", false, "don't create a unit test file")
	hookCmd.Flags().BoolVar(&models.CmdOptions.OnlyDryRun, "dry-run", false, "print only the changes, but don't write to disk")

	hookCmd.Flags().SortFlags = false
}
