package cmd

import (
	"fmt"

	"github.com/jaredbancroft/aoc2020/pkg/expense"
	"github.com/jaredbancroft/aoc2020/pkg/helpers"
	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Advent of Code 2020 - Day 1",
	Long: `
	Advent of Code 2020 - Day 1
	
	Before you leave, the Elves in accounting just need you to fix 
	your expense report (your puzzle input); apparently, something 
	isn't quite adding up.

	Specifically, they need you to find the two entries that 
	sum to 2020 and then multiply those two numbers together.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		expenses, err := helpers.ReadIntFile(input)
		if err != nil {
			return err
		}
		report := expense.NewReport(expenses)
		val := report.FindEntries(2020)
		fmt.Println("Product is: ", val)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
