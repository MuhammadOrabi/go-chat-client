package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// BaseURL ...
var BaseURL string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Clinet is a manager for chat.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if os.Getenv("BASE_URL") != "" {
		BaseURL = os.Getenv("BASE_URL")
	} else {
		BaseURL = "http://0.0.0.0:3000"
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
