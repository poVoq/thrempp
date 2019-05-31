package cmd

import (
	"os"

	"github.com/bdlm/log"
	"github.com/spf13/cobra"
)

var (
	timestamps bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "thrempp",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}