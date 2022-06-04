package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Create a new React 'Hook'",
	Long:  "Create a new React 'Hook'.",
	Run: func(cmd *cobra.Command, args []string) {
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		// TODO: Add logic here

		fmt.Println("'Hook' called with dry-run?", dryRun)
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
	hookCmd.Flags().Bool("dry-run", false, "Print only the changes, but don't write to disk")
}
