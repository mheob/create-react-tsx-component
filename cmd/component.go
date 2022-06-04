package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Create a new React 'Component'",
	Long:  "Create a new React 'Component'.",
	Run: func(cmd *cobra.Command, args []string) {
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		// TODO: Add logic here

		fmt.Println("'Component' called with dry-run?", dryRun)
	},
}

func init() {
	rootCmd.AddCommand(componentCmd)
	componentCmd.Flags().Bool("dry-run", false, "Print only the changes, but don't write to disk")
	pageCmd.Flags().BoolP("no-storybook", "s", false, "don't create a storybook stories file")
	pageCmd.Flags().BoolP("no-test", "t", false, "don't create a unit test file")
}
