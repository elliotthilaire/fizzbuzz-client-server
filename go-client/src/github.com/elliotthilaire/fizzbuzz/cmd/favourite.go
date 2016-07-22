package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/elliotthilaire/fizzbuzz/fizzbuzzutil"
)

// favouriteCmd represents the favourite command
var favouriteCmd = &cobra.Command{
	Use:   "favourite NUMBER",
	Short: "mark a number as a favourite",
	Long: `Marks a fizzbuzz results with cake
to show that you like it and that,
it is in fact, not a lie.`,

	Run: func(cmd *cobra.Command, args []string) {

    fmt.Println("adding favourites to ", host)

    for _,number := range args {
   	  fizzbuzzutil.AddFavourite(number)
    }
	},
}

func init() {
	RootCmd.AddCommand(favouriteCmd)
}
