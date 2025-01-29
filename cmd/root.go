package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sCodingCh",
	Short: "S Coding Challenge",
	Long:  "S Coding Challenge",
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Unexpected exection error, err %v", err)
	}
}
