package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/elliotthilaire/fizzbuzz/fizzbuzzutil"
)

// removeCmd represents the favourite command
var unfavouriteCmd = &cobra.Command{
  Use:   "unfavourite NUMBER",
  Short: "remove numbers from your favourites",
  Long: `Removes the cake from a number`,

  Run: func(cmd *cobra.Command, args []string) {

    fmt.Println("favourite called", host)

    for _,number := range args {
      fizzbuzzutil.DeleteFavourite(number)
    }
  },
}

func init() {
  RootCmd.AddCommand(unfavouriteCmd)
}
