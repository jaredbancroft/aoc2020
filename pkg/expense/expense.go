package expense

import "errors"

// Report is an integer slice of your expenses
type Report struct {
	expenses []int
}

//NewReport creates a new expense report
func NewReport(expenses []int) *Report {
	return &Report{expenses: expenses}
}

//Find2Entries will find the two entries which sum to the target
//and return their product
func (r *Report) Find2Entries(target int) (int, error) {
	hashmap := make(map[int]int)

	for _, expense := range r.expenses {
		if val, ok := hashmap[expense]; ok {
			return expense * val, nil
		}

		hashmap[target-expense] = expense
	}

	return -1, errors.New("No suitable entries")
}

//Find3Entries will find the 3 entries which sum to the target
//and return their produce
func (r *Report) Find3Entries(target int) (int, error) {

	for _, expense := range r.expenses {
		subset, err := r.Find2Entries(target - expense)
		if err != nil {
			continue
		}
		return expense * subset, nil
	}

	return -1, errors.New("No suitable entries")
}
