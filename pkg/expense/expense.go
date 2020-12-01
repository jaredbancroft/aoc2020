package expense

// Report is an integer slice of your expenses
type Report struct {
	expenses []int
}

//NewReport creates a new expense report
func NewReport(expenses []int) *Report {
	return &Report{expenses: expenses}
}

//FindEntries will find the two entries which sum to the target
func (r *Report) FindEntries(target int) int {
	hashmap := make(map[int]int)

	for _, expense := range r.expenses {
		if val, ok := hashmap[expense]; ok {
			return expense * val
		}

		hashmap[target-expense] = expense
	}
	return 0
}
