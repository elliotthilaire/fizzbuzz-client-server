package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var page int
var per_page int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [--page N] [--per_page N]",
	Short: "list a page of of fizzbuzz results",
	Long: `
Fetches a page of fizzbuzz results from the fizzbuzz server.
Displays favourite results with cake.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("list called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().IntVar(&page, "page", 1, "page number of results")
	listCmd.PersistentFlags().IntVar(&per_page, "per_page", 100, "number of results to show per page")
}
