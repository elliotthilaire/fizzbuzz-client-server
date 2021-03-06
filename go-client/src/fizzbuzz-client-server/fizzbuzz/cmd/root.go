package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var host string
var port string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "fizzbuzz",
	Short: "Interact with a fizzbuzz server",
	Long: `fizzbuzz is a CLI client for a fizzbuzz server.
You can list pages of fizzbuzz results and mark your favourites.

Favourites are marked with cake to show they are not a lie.

Examples:
fizzbuzz list --page 2 --per-page 10
fizzbuzz favourite 1 2 3 7 11
fizzbuzz unfavourite 1 11`,
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

	RootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "fizzbuzz server host")
	RootCmd.PersistentFlags().StringVar(&port, "port", "3000", "fizzbuzz server port")
}
