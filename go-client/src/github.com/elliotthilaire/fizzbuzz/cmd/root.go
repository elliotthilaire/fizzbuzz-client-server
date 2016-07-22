package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var host string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "fizzbuzz",
	Short: "Interact with a fizzbuzz server",
	Long: `fizzbuzz is a CLI client for a fizzbuzz server.
You can list pages of fizzbuzz results and also mark your favourites.

It is written in Go using the Cobra CLI library.`,
// Uncomment the following line if your bare application
// has an action associated with it:
//	Run: func(cmd *cobra.Command, args []string) { },

}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&host, "host", "localhost:3000", "fizzbuzz server")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

