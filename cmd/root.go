package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var input string

var rootCmd = &cobra.Command{
	Use:   "aoc2020",
	Short: "Advent of Code 2020",
	Long:  `Saving Christmas with code`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableSuggestions = true
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Input file (required)")
	rootCmd.MarkPersistentFlagRequired("input")
}
