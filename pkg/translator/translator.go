package translator

//Rule is a translator rule
type Rule struct {
	name                         string
	r1Low, r1High, r2Low, r2High int
}

//Ruler has some rules
type Ruler struct {
	rules []Rule
}

//NewRule makes a new rule
func NewRule(name string, low1, high1, low2, high2 int) Rule {
	return Rule{
		name:   name,
		r1Low:  low1,
		r1High: high1,
		r2Low:  low2,
		r2High: high2,
	}
}

//NewRuler makes a new ruler
func NewRuler(rules []Rule) Ruler {
	return Ruler{rules: rules}
}

//Check to see if the rules match
func (r Rule) Check(number int) bool {
	match := false
	if (number >= r.r1Low && number <= r.r1High) || (number >= r.r2Low && number <= r.r2High) {
		match = true
	}
	return match
}

//CheckAll checks all the rules against all the values
func (rr Ruler) CheckAll(myTicket []int, tickets [][]int) (int, int) {
	total := 0
	validTickets := make([][]int, 0)
	for _, ticket := range tickets {
		isValid := true
		for _, number := range ticket {
			atLeastOneValid := false
			for _, rule := range rr.rules {
				if rule.Check(number) == true {
					atLeastOneValid = true
					break
				}
			}
			if atLeastOneValid == false {
				isValid = false
				total += number
			}
		}
		if isValid {
			validTickets = append(validTickets, ticket)
		}
	}

	matches := make(map[string][]int)
	for _, rule := range rr.rules {
		for i := 0; i < len(validTickets[0]); i++ {
			numValid := 0
			for _, ticket := range validTickets {
				if rule.Check(ticket[i]) == true {
					numValid++
				}
			}

			if numValid == len(validTickets) {
				matches[rule.name] = append(matches[rule.name], i)
			}
		}
	}

	final := make(map[string]int)

	for {
		tmp := ""
		allDone := 0

		for k, match := range matches {
			if len(match) == 1 {
				final[k] = match[0]
				tmp = k
			}
		}

		for k, match := range matches {
			if len(match) > 0 {
				match = remove(match, final[tmp])
				matches[k] = match
			}
		}

		for _, match := range matches {
			if len(match) == 0 {
				allDone++
			}
		}

		if len(matches) == allDone {
			break
		}
	}

	part2 := myTicket[final["departure date"]] *
		myTicket[final["departure location"]] *
		myTicket[final["departure platform"]] *
		myTicket[final["departure station"]] *
		myTicket[final["departure time"]] *
		myTicket[final["departure track"]]

	return total, part2
}

func remove(s []int, v int) []int {

	i := 0

	for _, n := range s {
		if n != v {
			s[i] = n
			i++
		}
	}
	s = s[:i]
	return s
}
