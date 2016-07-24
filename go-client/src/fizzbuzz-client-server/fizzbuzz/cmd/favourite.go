package cmd

import (
	"github.com/spf13/cobra"
	"fizzbuzz-client-server/fizzbuzz/fizzbuzzutil"
)

// favouriteCmd represents the favourite command
var favouriteCmd = &cobra.Command{
	Use:   "favourite NUMBER [NUMBER..]",
	Short: "mark numbers as favourite",
	Long: `Adds cake to numbers passed in as arguments`,

	Run: func(cmd *cobra.Command, args []string) {

    if len(args) == 0 {
      cmd.Help()
      return
    }

    fizzbuzz := fizzbuzzutil.Api{Host: host, Port: port}

    for _,number := range args {
   	  fizzbuzz.AddFavourite(number)
    }
	},
}

func init() {
	RootCmd.AddCommand(favouriteCmd)
}
