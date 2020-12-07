package cmd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jaredbancroft/aoc2020/pkg/helpers"
	"github.com/jaredbancroft/aoc2020/pkg/luggage"
	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Advent of Code - 2020",
	Long: `--- Day 7: Handy Haversacks ---
	You land at the regional airport in time for your next flight. In fact, it looks like you'll even have time to grab some food: all flights are currently delayed due to issues in luggage processing.
	
	Due to recent aviation regulations, many rules (your puzzle input) are being enforced about bags and their contents; bags must be color-coded and must contain specific quantities of other color-coded bags. Apparently, nobody responsible for these regulations considered how long they would take to enforce!
	
	For example, consider the following rules:
	
	light red bags contain 1 bright white bag, 2 muted yellow bags.
	dark orange bags contain 3 bright white bags, 4 muted yellow bags.
	bright white bags contain 1 shiny gold bag.
	muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
	shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
	dark olive bags contain 3 faded blue bags, 4 dotted black bags.
	vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
	faded blue bags contain no other bags.
	dotted black bags contain no other bags.
	These rules specify the required contents for 9 bag types. In this example, every faded blue bag is empty, every vibrant plum bag contains 11 bags (5 faded blue and 6 dotted black), and so on.
	
	You have a shiny gold bag. If you wanted to carry it in at least one other bag, how many different bag colors would be valid for the outermost bag? (In other words: how many colors can, eventually, contain at least one shiny gold bag?)
	
	In the above rules, the following options would be available to you:
	
	A bright white bag, which can hold your shiny gold bag directly.
	A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
	A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
	A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
	So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.
	
	How many bag colors can eventually contain at least one shiny gold bag? (The list of rules is quite long; make sure you get all of it.)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		rules, err := helpers.ReadStringFile(input)
		if err != nil {
			return err
		}
		var baseLuggage [][]string
		var luggageList []luggage.Bag
		re1 := regexp.MustCompile(`(\w+\s\w+)\sbags\scontain`)
		re2 := regexp.MustCompile(`contain\s(?:no\sother\sbags\.|(?:(\d)\s(\w+\s\w+)))`)
		re3 := regexp.MustCompile(`,\s(\d)\s(\w+\s\w+)`)
		bc := make(map[string]int)
		for _, rule := range rules {
			baseLuggage = re1.FindAllStringSubmatch(rule, -1)
			bag := luggage.NewBag(baseLuggage[0][1])
			contains := re2.FindAllStringSubmatch(rule, -1)
			name := contains[0][2]
			i, _ := strconv.Atoi(contains[0][1])
			bc[name] = i
			containsmore := re3.FindAllStringSubmatch(rule, -1)
			for _, cm := range containsmore {
				j, _ := strconv.Atoi(cm[1])
				bc[cm[2]] = j
			}
			bag.Contains = bc

			luggageList = append(luggageList, *bag)
			bc = make(map[string]int)
		}

		simplify := []string{"shiny gold"}
		simplify2 := []string{}
		res := []luggage.Bag{}
		for {

			var before, after int

			for _, s := range simplify {
				for i, ll := range luggageList {
					if _, ok := ll.Contains[s]; ok {
						simplify2 = append(simplify2, ll.Name)
						res = append(res, ll)
						luggageList[i] = luggage.Bag{}
					}
				}
			}
			simplify = simplify2
			simplify2 = []string{}
			after = len(simplify)

			if before == after {
				break
			}
		}
		fmt.Println(len(res))
		var bag luggage.Bag
		for _, ll := range luggageList {
			if ll.Name == "shiny gold" {
				bag = ll
			}
		}

		fmt.Println(bag.Walk(luggageList))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)
}
