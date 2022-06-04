package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Create a new React 'Page'",
	Long:  "Create a new React 'Page'.",
	Run: func(cmd *cobra.Command, args []string) {
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		// TODO: Add logic here

		fmt.Println("'Page' called with dry-run?", dryRun)
	},
}

func init() {
	rootCmd.AddCommand(pageCmd)
	pageCmd.Flags().Bool("dry-run", false, "print only the changes, but don't write to disk")
}
