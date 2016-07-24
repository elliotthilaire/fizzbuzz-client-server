package cmd

import (
  "github.com/spf13/cobra"
  "fizzbuzz-client-server/fizzbuzz/fizzbuzzutil"
)

// unfavouriteCmd represents the unfavourite command
var unfavouriteCmd = &cobra.Command{
  Use:   "unfavourite NUMBER",
  Short: "remove numbers from your favourites",
  Long: `Removes the cake from a number`,

  Run: func(cmd *cobra.Command, args []string) {
    fizzbuzz := fizzbuzzutil.Api{Host: host, Port: port}

    for _,number := range args {
      fizzbuzz.DeleteFavourite(number)
    }
  },
}

func init() {
  RootCmd.AddCommand(unfavouriteCmd)
}
