package cmd

import (
  "github.com/spf13/cobra"
  "fizzbuzz-client-server/fizzbuzz/fizzbuzzutil"
)

// unfavouriteCmd represents the unfavourite command
var unfavouriteCmd = &cobra.Command{
  Use:   "unfavourite NUMBER [NUMBER..]",
  Short: "unmark numbers as favourites",
  Long: `Removes cake from numbers passed in as arguments`,

  Run: func(cmd *cobra.Command, args []string) {

    if len(args) == 0 {
      cmd.Help()
      return
    }

    fizzbuzz := fizzbuzzutil.Api{Host: host, Port: port}

    for _,number := range args {
      fizzbuzz.DeleteFavourite(number)
    }
  },
}

func init() {
  RootCmd.AddCommand(unfavouriteCmd)
}
