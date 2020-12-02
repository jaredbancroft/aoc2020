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
	Long: `
	Advent of Code 2020

	After saving Christmas five years in a row, you've decided to take a vacation at a nice 
	resort on a tropical island. Surely, Christmas will go on without you.

	The tropical island has its own currency and is entirely cash-only. 
	The gold coins used there have a little picture of a starfish; the locals just call them stars. 
	None of the currency exchanges seem to have heard of them, but somehow, you'll need to find 
	fifty of these coins by the time you arrive so you can pay the deposit on your room.
	
	To save your vacation, you need to get all fifty stars by December 25th.
	
	Collect stars by solving puzzles. Two puzzles will be made available on each day 
	in the Advent calendar; the second puzzle is unlocked when you complete the first. 
	Each puzzle grants one star. Good luck!`,
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
