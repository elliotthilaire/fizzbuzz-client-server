package cmd

import (
	"github.com/spf13/cobra"
	"fizzbuzz-client-server/fizzbuzz/fizzbuzzutil"
)

// favouriteCmd represents the favourite command
var favouriteCmd = &cobra.Command{
	Use:   "favourite NUMBER",
	Short: "mark a number as a favourite",
	Long: `Marks a fizzbuzz results with cake
to show that you like it and that,
it is in fact, not a lie.`,

	Run: func(cmd *cobra.Command, args []string) {
    fizzbuzz := fizzbuzzutil.Api{Host: host, Port: port}

    for _,number := range args {
   	  fizzbuzz.AddFavourite(number)
    }
	},
}

func init() {
	RootCmd.AddCommand(favouriteCmd)
}
