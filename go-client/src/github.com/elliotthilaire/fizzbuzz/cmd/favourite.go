package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// favouriteCmd represents the favourite command
var favouriteCmd = &cobra.Command{
	Use:   "favourite NUMBER",
	Short: "mark a number as a favourite",
	Long: `Marks a fizzbuzz results with cake
to show that you like it,
and that is in fact, not a lie.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("favourite called", host)
	},
}

func init() {
	RootCmd.AddCommand(favouriteCmd)

	favouriteCmd.Flags().BoolP("delete", "d", false, "Deletes the number as a favourite")
}
